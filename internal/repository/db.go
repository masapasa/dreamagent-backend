package repository

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masapasa/dreamagent-backend/internal/config"
	"github.com/masapasa/dreamagent-backend/internal/models"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewPostgresDB(cfg *config.Config) *DB {
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	db := &DB{Pool: pool}
	db.initTables()
	return db
}

func (db *DB) initTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS agents (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id TEXT NOT NULL,
			name TEXT NOT NULL,
			description TEXT,
			model TEXT NOT NULL DEFAULT 'claude-4',
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS memories (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			agent_id UUID REFERENCES agents(id) ON DELETE CASCADE,
			key TEXT NOT NULL,
			value TEXT NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS sessions (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			agent_id UUID REFERENCES agents(id) ON DELETE CASCADE,
			input JSONB NOT NULL,
			output JSONB,
			status TEXT NOT NULL DEFAULT 'pending',
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS subscriptions (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id TEXT NOT NULL,
			stripe_customer_id TEXT,
			stripe_subscription_id TEXT UNIQUE,
			status TEXT NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,
	}
	for _, q := range queries {
		_, err := db.Pool.Exec(context.Background(), q)
		if err != nil {
			log.Printf("Table init error: %v", err)
		}
	}
}

func (db *DB) CreateAgent(ctx context.Context, agent models.Agent) (models.Agent, error) {
	query := `INSERT INTO agents (user_id, name, description, model) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := db.Pool.QueryRow(ctx, query, agent.UserID, agent.Name, agent.Description, agent.Model).Scan(&agent.ID, &agent.CreatedAt)
	return agent, err
}

func (db *DB) GetAgentsByUser(ctx context.Context, userID string) ([]models.Agent, error) {
	rows, err := db.Pool.Query(ctx, `SELECT id, user_id, name, description, model, created_at FROM agents WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var agents []models.Agent
	for rows.Next() {
		var a models.Agent
		if err := rows.Scan(&a.ID, &a.UserID, &a.Name, &a.Description, &a.Model, &a.CreatedAt); err != nil {
			return nil, err
		}
		agents = append(agents, a)
	}
	return agents, nil
}

func (db *DB) DreamAgent(ctx context.Context, agentID uuid.UUID) error {
	// Production: query past sessions + memories, call Anthropic to analyze/refine, upsert cleaned memories
	log.Printf("[Dream] Refining memories and knowledge base for agent %s (inspired by Anthropic Dreaming)", agentID)
	return nil
}
