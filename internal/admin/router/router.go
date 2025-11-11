package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"goadmin/internal/admin/controller"
	"goadmin/internal/admin/repository"
	"goadmin/internal/admin/service"
	appConfig "goadmin/pkg/config"
)

// Register wires up all admin routes on the provided engine.
func Register(engine *gin.Engine, db *gorm.DB, cfg *appConfig.Config) {
	operatorRepo := repository.NewOperatorRepository(db)
	authService := service.NewAuthService(operatorRepo, cfg)
	authController := controller.NewAuthController(authService)

	adminGroup := engine.Group("/admin")
	{
		adminGroup.POST("/account/login", authController.Login)
	}
}

// NewEngine builds a standalone Gin engine with the admin routes, kept for compatibility.
func NewEngine(db *gorm.DB, cfg *appConfig.Config) *gin.Engine {
	engine := gin.Default()
	Register(engine, db, cfg)
	return engine
}
