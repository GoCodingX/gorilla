TOOLS_PATH := $(CURDIR)/.tools
LINTER_BINARY := $(TOOLS_PATH)/golangci-lint

# install golangci-lint linter checks
.PHONY: install-linter
install-linter:
	@if [ ! -f $(LINTER_BINARY) ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLS_PATH) v1.63.4; \
	fi

# runs golangci-lint linter checks
.PHONY: lint
lint: install-linter
	$(LINTER_BINARY) run --fix

# recreate the vendors directory and tidy dependencies
.PHONY: mod-clean
mod-clean:
	@rm -rf vendor && go mod tidy && go mod vendor

.PHONY: generate
generate:
	go generate ./...

.PHONY: dev
dev: run-dependencies
	@go run cmd/quotes-api/main.go

.PHONY: run-dependencies
run-dependencies:
	docker compose up -d

.PHONY: test
test:
	@go test ./...
