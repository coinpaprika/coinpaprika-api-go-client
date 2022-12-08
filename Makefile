SHELL := /bin/bash
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

.PHONY: check format help test tidy

format: ## Format go code with goimports
	@go install golang.org/x/tools/cmd/goimports@latest
	@goimports -l -w .

test: ## Run tests
	@go test -race ./...

tidy: ## Run go mod tidy
	@go mod tidy

check: ## Linting and static analysis
	@if test ! -e ./bin/golangci-lint; then \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh; \
	fi

	@./bin/golangci-lint run -c .golangci.yml

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
