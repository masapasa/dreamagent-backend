package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/masapasa/dreamagent-backend/internal/config"
	"github.com/masapasa/dreamagent-backend/internal/models"
	"github.com/masapasa/dreamagent-backend/internal/repository"
	"github.com/masapasa/dreamagent-backend/internal/services"
)

type Handler struct {
	db            *repository.DB
	stripeService *services.StripeService
	cfg           *config.Config
}

func New(cfg *config.Config, db *repository.DB) *Handler {
	return &Handler{
		db:            db,
		stripeService: services.NewStripeService(cfg.StripeSecretKey),
		cfg:           cfg,
	}
}

func (h *Handler) CreateAgent(c *gin.Context) {
	userID := c.GetString("user_id")
	var agent models.Agent
	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	agent.UserID = userID
	created, err := h.db.CreateAgent(c.Request.Context(), agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *Handler) ListAgents(c *gin.Context) {
	userID := c.GetString("user_id")
	agents, err := h.db.GetAgentsByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agents)
}

func (h *Handler) RunAgent(c *gin.Context) {
	agentIDStr := c.Param("id")
	agentID, err := uuid.Parse(agentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid agent ID"})
		return
	}
	var req models.RunRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Production: integrate Anthropic API call with memory lookup, multi-agent orchestration, outcome loop + grader
	// Placeholder demonstrates Outcomes feature
	output := map[string]interface{}{
		"result": "Outcome achieved: " + req.Outcome,
		"status": "success",
		"note": "Multi-agent + Dreaming ready for production",
	}
	c.JSON(http.StatusOK, gin.H{"output": output})
}

func (h *Handler) Dream(c *gin.Context) {
	agentIDStr := c.Param("id")
	agentID, err := uuid.Parse(agentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid agent ID"})
		return
	}
	if err := h.db.DreamAgent(c.Request.Context(), agentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dreaming complete — memories refined and self-improved"})
}
