GO = go
GOPATH = $(shell $(GO) env GOPATH)
GOBIN = $(GOPATH)/bin

GREEN = \033[0;32m
NC = \033[0m

.PHONY: install-tools
install-tools:
	@if ! command -v $(GOBIN)/swag >/dev/null 2>&1; then \
		echo "[$(GREEN)*$(NC)] Installing swag"; \
		$(GO) install github.com/swaggo/swag/cmd/swag@latest; \
	fi

.PHONY: gen-docs
gen-docs: install-tools
	@echo "[$(GREEN)*$(NC)] Generating swagger docs"
	@$(GOBIN)/swag init

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
