# DreamAgent Backend

**The one AI product I want to build based on Anthropic’s 2026 Developer Conference info:**

**DreamAgent** — A fully managed, self-improving AI agent SaaS platform.

Agents get:
- **Dreaming**: Background refinement of memory stores from past sessions (merges duplicates, extracts patterns, updates institutional knowledge — exactly like Anthropic’s new feature).
- **Multi-agent orchestration**: Coordinator spawns parallel sub-agents.
- **Outcomes**: Goal-driven loops with built-in grader until success.

Backend powers user accounts, persistent agents, memory, sessions, billing.

## Tech Stack
- Go (production-ready, fast, scalable)
- Supabase Auth (JWT) + PostgreSQL
- Stripe Subscriptions
- Gin for API
- Ready for Docker deploy

## Quick Start
1. `cp .env.example .env` and fill keys
2. `go mod tidy`
3. `go run cmd/server/main.go`
4. Agents auto-persist memories & dream on demand.

Pushed via connected GitHub tools as requested.
