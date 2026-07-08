APP_NAME = ms-stream-api

.DEFAULT_GOAL := help

.PHONY: help init run docker-build docker-run docker-up docker-down

init: ## Install Go dependencies and create local logs directory
	go mod download
	go mod tidy
	mkdir -p ./logs/

run: ## Run the app locally
	go run main.go

docker-build: ## Build the Docker image
	docker build -t $(APP_NAME):latest .


docker-run: ## Run the Docker container locally
	docker run --rm \
		--add-host=localhost:host-gateway \
		--publish 8081:8081 \
		--volume=$(CURDIR)/assets:/app/assets \
		--volume=$(CURDIR)/logs:/app/logs \
		--name $(APP_NAME) \
		$(APP_NAME):latest

docker-up: ## Start the compose service
	docker compose up -d $(APP_NAME)

docker-down: ## Stop the compose service
	docker compose down

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
