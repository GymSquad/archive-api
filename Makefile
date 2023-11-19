GO = go
GOPATH = $(shell $(GO) env GOPATH)
GOBIN = $(GOPATH)/bin

GREEN = \033[0;32m
NC = \033[0m

.PHONY: install-tools
install-tools:
	@if ! command -v $(GOBIN)/oapi-codegen >/dev/null 2>&1; then \
		echo "[$(GREEN)*$(NC)] Installing oapi-codegen"; \
		$(GO) install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest; \
	fi

.PHONY: generate
generate: install-tools
	@echo "[$(GREEN)*$(NC)] Generating the server code from the OpenAPI spec"
	@$(GOBIN)/oapi-codegen -package api -generate types,server,spec docs/openapi.yml > api/api.gen.go

.PHONY: deps
deps:
	@echo "[$(GREEN)*$(NC)] Installing dependencies"
	@$(GO) mod download

.PHONY: run
run:
	@echo "[$(GREEN)*$(NC)] Running server"
	@$(GO) run main.go

.PHONY: build
build:
	@echo "[$(GREEN)*$(NC)] Building server"
	@$(GO) build -o bin/server main.go
