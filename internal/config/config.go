package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseURL         string
	SupabaseAnonKey     string
	JWTSecret           string
	StripeSecretKey     string
	StripeWebhookSecret string
	AnthropicAPIKey     string
	DatabaseURL         string
	Port                string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: no .env file")
	}
	return &Config{
		SupabaseURL:         os.Getenv("SUPABASE_URL"),
		SupabaseAnonKey:     os.Getenv("SUPABASE_ANON_KEY"),
		JWTSecret:           os.Getenv("SUPABASE_JWT_SECRET"),
		StripeSecretKey:     os.Getenv("STRIPE_SECRET_KEY"),
		StripeWebhookSecret: os.Getenv("STRIPE_WEBHOOK_SECRET"),
		AnthropicAPIKey:     os.Getenv("ANTHROPIC_API_KEY"),
		DatabaseURL:         os.Getenv("DATABASE_URL"),
		Port:                os.Getenv("PORT"),
	}
}
