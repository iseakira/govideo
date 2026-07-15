APP_NAME = ms-recommendation-api
EXTERNAL_APPS = postgresql

.PHONY: help run docker-build docker-build-no-cache docker-run external-up external-down docker-up docker-down

run: ## Run the app locally
	python main.py

docker-build: ## Build the Docker image
	docker build -t $(APP_NAME) .

docker-build-no-cache: ## Build the Docker image without cache
	docker build --no-cache -t $(APP_NAME) .

docker-run: ## Run the Docker container locally
	docker run --rm \
		--add-host=localhost:host-gateway \
		--publish 8082:8082 \
		--volume=logs:/app/logs \
		--name $(APP_NAME) \
		$(APP_NAME):latest

external-up: ## Start PostgreSQL for local development
	docker compose up -d $(EXTERNAL_APPS)

external-down: ## Stop PostgreSQL for local development
	docker compose down

docker-up: ## Start the compose service
	docker compose up -d $(APP_NAME)

docker-down: ## Stop the compose stack
	docker compose down

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'