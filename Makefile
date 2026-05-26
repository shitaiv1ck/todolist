include .env
export

env-up:
	@docker compose up -d todolist-postgres

env-down:
	@docker compose down todolist-postgres

migrate-up:
	@migrate -path ./migrations -database \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
	up

migrate-down:
	@migrate -path ./migrations -database \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
	down

migrate-force:
	@migrate -path ./migrations -database \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
	force 1

app-run:
	@go run ./cmd/todolist/main.go

	