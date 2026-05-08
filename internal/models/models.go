package models

import (
	"time"

	"github.com/google/uuid"
)

type Agent struct {
	ID          uuid.UUID `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Model       string    `json:"model"`
	CreatedAt   time.Time `json:"created_at"`
}

type Memory struct {
	ID        uuid.UUID `json:"id"`
	AgentID   uuid.UUID `json:"agent_id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"` // markdown content
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	ID        uuid.UUID              `json:"id"`
	AgentID   uuid.UUID              `json:"agent_id"`
	Input     map[string]interface{} `json:"input"`
	Output    map[string]interface{} `json:"output"`
	Status    string                 `json:"status"`
	CreatedAt time.Time              `json:"created_at"`
}

type RunRequest struct {
	Outcome string                 `json:"outcome"`
	Input   map[string]interface{} `json:"input"`
}
