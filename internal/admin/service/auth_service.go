package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"goadmin/internal/admin/dto"
	"goadmin/internal/admin/repository"
	"goadmin/internal/common/model"
	appConfig "goadmin/pkg/config"
	"goadmin/pkg/utils"
)

var (
	// ErrInvalidCredentials indicates account or password mismatch.
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrAccountDisabled indicates the operator is frozen.
	ErrAccountDisabled = errors.New("account is disabled")
)

// AuthService provides authentication use cases.
type AuthService struct {
	repo        repository.OperatorRepository
	tokenSecret string
	tokenTTL    time.Duration
}

// NewAuthService creates a new AuthService instance.
func NewAuthService(repo repository.OperatorRepository, cfg *appConfig.Config) *AuthService {
	return &AuthService{
		repo:        repo,
		tokenSecret: cfg.Auth.TokenSecret,
		tokenTTL:    cfg.Auth.TokenTTL,
	}
}

// Login validates the credentials and returns an authentication token.
func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest, clientIP string) (*dto.LoginResponse, error) {
	account := utils.SafeString(req.Username)
	password := utils.SafeString(req.Password)
	if account == "" || password == "" {
		return nil, ErrInvalidCredentials
	}

	operator, err := s.repo.GetByAccount(ctx, account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if operator.Status != 1 {
		return nil, ErrAccountDisabled
	}

	if !s.verifyPassword(operator, password) {
		return nil, ErrInvalidCredentials
	}

	now := time.Now()
	if err := s.repo.UpdateLoginMeta(ctx, operator.ID, clientIP, now); err != nil {
		return nil, err
	}

	roles := s.parseRoles(operator.Roles)
	token, err := s.generateToken(operator, roles, now)
	if err != nil {
		return nil, fmt.Errorf("generate token: %w", err)
	}

	return &dto.LoginResponse{
		Token:     token,
		ExpiresAt: now.Add(s.tokenTTL),
		OperatorInfo: dto.OperatorInfo{
			ID:       operator.ID,
			Name:     operator.CName,
			Phone:    operator.Phone,
			Email:    operator.Email,
			Nickname: operator.Nickname,
			Roles:    roles,
			IsAdmin:  operator.IsAdmin == "1",
		},
	}, nil
}

func (s *AuthService) verifyPassword(operator *model.Operator, raw string) bool {

	salt := appConfig.GetString("OPERATOR_PASSWORD_SALT", "default_salt")
	raw = fmt.Sprintf("%s%s%s", raw, salt, operator.CreateTime.Format("2006-01-02 15:04:05"))

	hashed := utils.MD5String(utils.MD5String(raw))
	fmt.Println(operator.Password)
	fmt.Println(hashed)
	return strings.EqualFold(operator.Password, hashed)
}

func (s *AuthService) parseRoles(roles string) []string {
	trimmed := utils.SafeString(roles)
	if trimmed == "" {
		return nil
	}
	parts := strings.Split(trimmed, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if p := utils.SafeString(part); p != "" {
			result = append(result, p)
		}
	}
	return result
}

func (s *AuthService) generateToken(operator *model.Operator, roles []string, now time.Time) (string, error) {
	expiresAt := now.Add(s.tokenTTL)
	isAdmin := operator.IsAdmin == "1"

	claims := jwt.MapClaims{
		"sub":      fmt.Sprintf("%d", operator.ID),
		"name":     operator.CName,
		"roles":    roles,
		"is_admin": isAdmin,
		"jti":      uuid.NewString(),
		"iat":      now.Unix(),
		"nbf":      now.Unix(),
		"exp":      expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.tokenSecret))
}
