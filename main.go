package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"goadmin/internal/admin/router"
	appConfig "goadmin/pkg/config"
	"goadmin/pkg/database"
)

func main() {
	log.SetOutput(os.Stdout)

	// Load environment variables from .env if present; ignore error when file missing.
	_ = godotenv.Load()

	cfg, err := appConfig.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	db, err := database.NewMySQL(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	router.Register(engine, db, cfg)

	if err := engine.Run(cfg.Server.Address); err != nil {
		log.Fatalf("启动 HTTP 服务失败: %v", err)
	}
}
