APP_NAME = ms-api
EXTERNAL_APPS = postgresql
AUTH_ENABLE ?= false
.PHONY: help external-up external-down external-clean docker-build docker-build-no-cache docker-run

external-up: ## Start PostgreSQL for local development
	docker compose up -d $(EXTERNAL_APPS)

external-down: ## Stop PostgreSQL for local development
	docker compose down

external-clean: ## Stop PostgreSQL and DELETE all data (volumes)
	docker compose down -v

docker-build: ## Build the Docker image
	docker build -t $(APP_NAME) .

docker-build-no-cache: ## Build the Docker image without cache
	docker build --no-cache -t $(APP_NAME) .

docker-run: ## Run the Docker container locally
	docker run --rm \
		--add-host=localhost:host-gateway \
		--publish 8080:8080 \
		--env AUTH_ENABLE=$(AUTH_ENABLE) \
		--volume=assets:/go/src/$(APP_NAME)/assets \
		--volume=logs:/go/src/$(APP_NAME)/logs \
		--name $(APP_NAME) \
		$(APP_NAME):latest

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'