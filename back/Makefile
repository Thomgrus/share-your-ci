go-dependencies: ## Give dependencies
	@go get -d -v ./...

go-run: ## Run directly with go
	@go run main.go

build: ## Build docker image
	@docker-compose build core-api

run: ## Run api project
	@docker-compose up core-api

core-api: ## Run api project in detach mode
	@docker-compose up -d core-api

stop: ## Stop project running with docker-compose
	@docker-compose down

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help