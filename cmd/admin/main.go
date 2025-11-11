package main

import (
	"log"
	"os"

	"goadmin/internal/admin/router"
	appConfig "goadmin/pkg/config"
	"goadmin/pkg/database"
)

func main() {
	log.SetOutput(os.Stdout)

	cfg, err := appConfig.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	db, err := database.NewMySQL(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	engine := router.NewEngine(db, cfg)

	if err := engine.Run(cfg.Server.Address); err != nil {
		log.Fatalf("启动 HTTP 服务失败: %v", err)
	}
}
