TOOLS_PATH := $(CURDIR)/.tools
LINTER_BINARY := $(TOOLS_PATH)/golangci-lint

# install golangci-lint linter checks
install-linter:
	@if [ ! -f $(LINTER_BINARY) ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLS_PATH) v1.63.4; \
	fi
.PHONY: install-linter

# runs golangci-lint linter checks
lint: install-linter
	$(LINTER_BINARY) run --fix
.PHONY: lint

# recreate the vendors directory and tidy dependencies
mod-clean:
	@rm -rf vendor && go mod tidy && go mod vendor
.PHONY: mod-clean

generate:
	go generate ./...
.PHONY: generate

dev:
	@go run cmd/quotes-api/main.go
.PHONY: dev

test:
	@go test ./...
.PHONY: test
