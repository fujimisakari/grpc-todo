GIT_REF := $(shell git describe --always --tag)
VERSION ?= commit-$(GIT_REF)

PROTO_PATH := ./app/driver/proto
PD_PATH := ./app/adapter/pb

export GOBIN := $(PWD)/bin

.PHONY: help
help: ## show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[\/a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

## run server
.PHONY: dev
dev: ## dev
	@go run cmd/main.go

## build
build/server: ## server build
	CGO_ENABLED=0 go build -o bin/server \
        -ldflags "-X main.version=$(VERSION)" \
        github.com/fujimisakari/grpc-todo/cmd

.PHONY: test
test: ## test
	@echo $(GOOGLEAPIS_PATH)

## tools
.PHONY: install-tools
install-tools: ## install tools
	@./tools/installer.sh

## formatter
.PHONY: fmt/goimports
fmt/goimports:  ## fmt by goimports
	@find . -type f -iname "*.go" -not -path "./vendor/**" -not -path "*/docs/*" -not -name "*.yo.go" |\
	xargs -I "{}" $(GOBIN)/goimports-reviser -project-name github.com/fujimisakari/grpc-todo "{}"

## lint check
.PHONY: lint
lint: ## run lint
	@$(GOBIN)/golangci-lint run --out-format=line-number ./...
	@make fmt/goimports

.PHONY: ensure-tidy
ensure-tidy: ## check that go mod tidy have already done
	@go mod tidy
	@cd tools && go mod tidy
	@git diff --exit-code go.mod go.sum tools/go.mod tools/go.sum

## proto compile
proto/compile: proto/getting-googleapis ## proto compile
	@$(eval GOOGLEAPIS_PATH := $(shell go list -m -json github.com/googleapis/googleapis | jq -r .Dir))
	@echo "Using googleapis path: $(GOOGLEAPIS_PATH)"
	@echo "Compiling proto files..."
	@protoc --proto_path=$(GOOGLEAPIS_PATH) \
		--proto_path=$(PROTO_PATH) \
		--go_out=$(PD_PATH) --go_opt=paths=source_relative \
		--go-grpc_out=$(PD_PATH) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(PD_PATH) --grpc-gateway_opt=paths=source_relative \
		$(PROTO_PATH)/todo.proto
	@echo "Proto compilation completed successfully"
	@go mod tidy

proto/getting-googleapis: ## getting googleapis
	@echo "Getting googleapis dependency..."
	@go get github.com/googleapis/googleapis

