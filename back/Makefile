redis-start: ## Start Redis
	@docker pull redis
	@docker container run --name redis-test-instance -p 6379:6379 -d redis

redis-stop: ## Start Redis
	@docker container stop redis-test-instance

go-dependencies: ## Give dependencies
	@go get -d -v ./...

go-run: ## Run directly with go
	@go run .

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help