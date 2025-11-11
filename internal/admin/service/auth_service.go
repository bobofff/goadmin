package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

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

	token := s.generateToken(operator.ID, now)

	return &dto.LoginResponse{
		Token:     token,
		ExpiresAt: now.Add(s.tokenTTL),
		OperatorInfo: dto.OperatorInfo{
			ID:       operator.ID,
			Name:     operator.CName,
			Phone:    operator.Phone,
			Email:    operator.Email,
			Nickname: operator.Nickname,
			Roles:    s.parseRoles(operator.Roles),
			IsAdmin:  operator.IsAdmin == "1",
		},
	}, nil
}

func (s *AuthService) verifyPassword(operator *model.Operator, raw string) bool {
	hashed := utils.MD5String(raw)
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

func (s *AuthService) generateToken(operatorID int, now time.Time) string {
	// NOTE: Placeholder token implementation. Replace with JWT or other secure mechanism if needed.
	return uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(s.tokenSecret+now.String())).String()
}
