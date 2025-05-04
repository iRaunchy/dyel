package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iraunchy/dyel/db"
	"github.com/iraunchy/dyel/internal/config"
	"github.com/iraunchy/dyel/internal/handlers"
	"github.com/iraunchy/dyel/internal/repos"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	dbConn, err := db.Init(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database init error: %v", err)
	}

	repo := repos.NewGORMProgramRepo(dbConn)
	h := handlers.NewHandler(repo)

	router := gin.Default()
	h.RegisterRoutes(router)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
