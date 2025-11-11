package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goadmin/internal/admin/dto"
	"goadmin/internal/admin/service"
	"goadmin/pkg/response"
)

// AuthController handles authentication endpoints.
type AuthController struct {
	authService *service.AuthService
}

// NewAuthController constructs a new AuthController.
func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login authenticates an operator and returns a token.
func (ctl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, 4001, "请求参数不正确")
		return
	}

	result, err := ctl.authService.Login(c.Request.Context(), req, c.ClientIP())
	if err != nil {
		switch err {
		case service.ErrInvalidCredentials:
			response.Error(c, http.StatusUnauthorized, 4010, "账号或密码错误")
		case service.ErrAccountDisabled:
			response.Error(c, http.StatusForbidden, 4030, "账号已被禁用")
		default:
			response.Error(c, http.StatusInternalServerError, 5000, "登录失败，请稍后再试")
		}
		return
	}

	response.Success(c, result)
}
