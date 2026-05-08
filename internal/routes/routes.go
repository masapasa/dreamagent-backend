package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/masapasa/dreamagent-backend/internal/config"
	"github.com/masapasa/dreamagent-backend/internal/handlers"
	"github.com/masapasa/dreamagent-backend/internal/middleware"
)

func Setup(r *gin.Engine, h *handlers.Handler, cfg *config.Config) {
	// Protected routes require Supabase auth
	r.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	api := r.Group("/api")
	{
		api.POST("/agents", h.CreateAgent)
		api.GET("/agents", h.ListAgents)
		api.POST("/agents/:id/run", h.RunAgent)
		api.POST("/agents/:id/dream", h.Dream)
	}

	// Public routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "product": "DreamAgent"})
	})

	// Stripe webhook (no auth)
	r.POST("/stripe/webhook", func(c *gin.Context) {
		// implement body + sig verification using h.stripeService
		c.JSON(200, gin.H{"status": "ok"})
	})
}
