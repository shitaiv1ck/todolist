include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d todolist-postgres

env-down:
	@docker compose down todolist-postgres

migrate-create:
	@if [ -z "$(seq)"]; then \
		echo "seq don't hava a value. pls, try again with seq=value"; \
		exit 1;\
	fi; \
	docker compose run --rm todolist-migrate \
	create -ext sql -dir ./migrations -seq "${seq}"


migrate-up:
	@docker compose run --rm todolist-migrate \
	-path ./migrations \
	-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todolist-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
	up 

migrate-down:
	@docker compose run --rm todolist-migrate \
	-path ./migrations \
	-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todolist-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
	down 

migrate-force:
	@docker compose run --rm todolist-migrate \
	-path ./migrations \
	-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todolist-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
	force 1

app-run:
	@go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todolist/main.go

app-deploy:
	@docker compose up -d --build todolist

app-stop:
	@docker compose down todolist

	