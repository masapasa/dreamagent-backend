package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/masapasa/dreamagent-backend/internal/config"
	"github.com/masapasa/dreamagent-backend/internal/handlers"
	"github.com/masapasa/dreamagent-backend/internal/repository"
	"github.com/masapasa/dreamagent-backend/internal/routes"
)

func main() {
	cfg := config.Load()
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	db := repository.NewPostgresDB(cfg)
	h := handlers.New(cfg, db)
	r := gin.Default()
	routes.Setup(r, h, cfg)

	log.Printf("🚀 DreamAgent backend listening on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
