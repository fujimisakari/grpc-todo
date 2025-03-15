-include .env
export

GOBIN := $(PWD)/bin

.PHONY: help
help: ## show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[\/a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

## run server
.PHONY: dev
dev: ## dev
	@go run cmd/main.go

## build
GIT_REF := $(shell git describe --always --tag)
VERSION ?= commit-$(GIT_REF)

build/server: ## server build
	CGO_ENABLED=0 go build -o bin/server \
        -ldflags "-X main.version=$(VERSION)" \
        github.com/fujimisakari/grpc-todo/cmd

.PHONY: test
test: ## test
	@echo $(GOOGLEAPIS_PATH)

## instal tools command
.PHONY: install-tools
install-tools: ## install tools command
	@./tools/cmd/installer.sh

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
	@cd tools/cmd && go mod tidy
	@git diff --exit-code go.mod go.sum tools/cmd/go.mod tools/cmd/go.sum

## proto compile
PROTO_PATH := ./app/driver/proto
PD_PATH := ./app/adapter/pb

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

## spanner-emulator
SPANNER_PROJECT := grpc-todo-spanner-emulator-project
SPANNER_INSTANCE := grpc-todo-spanner-emulator-instance
SPANNER_DATABASE := grpc-todo-spanner-emulator-db

SPANNER_EMULATOR_CONTAINER_NAME := grpc-todo-spanner-emulator
SPANNER_EMULATOR_RUNNING_PORT_CMD := docker port $(SPANNER_EMULATOR_CONTAINER_NAME) 9010/tcp 2> /dev/null | head -n 1 | rev | cut -d ":" -f1 | rev
ifneq ($(shell $(SPANNER_EMULATOR_RUNNING_PORT_CMD)),)
	SPANNER_EMULATOR_RUNNING_PORT := $(shell $(SPANNER_EMULATOR_RUNNING_PORT_CMD))
	SPANNER_EMULATOR_HOST := localhost:$(SPANNER_EMULATOR_RUNNING_PORT)
endif

.PHONY: spanner-emulator/start
spanner-emulator/start:  ## for local development
	@if [ -z "$(SPANNER_EMULATOR_RUNNING_PORT)" ]; then \
		docker run --rm -d --name $(SPANNER_EMULATOR_CONTAINER_NAME) -P gcr.io/cloud-spanner-emulator/emulator; \
		make spanner-emulator/setup; \
	fi

.PHONY: spanner-emulator/stop
spanner-emulator/stop:
	@if [ -n "$(SPANNER_EMULATOR_RUNNING_PORT)" ]; then \
		docker stop $(SPANNER_EMULATOR_CONTAINER_NAME); \
	fi

spanner-emulator/setup: spanner-emulator/createinstance wrench/create spanner-emulator/seed ## setup spanner emulator

spanner-emulator/createinstance:  ## create spanner instance
	@go run ./tools/spanner-emulator

.PHONY: spanner-emulator/cli
spanner-emulator/cli:
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) \
	$(GOBIN)/spanner-cli --project $(SPANNER_PROJECT) --instance $(SPANNER_INSTANCE) --database $(SPANNER_DATABASE)

spanner-emulator/seed:
	@make wrench/apply DML=db/spanner/seed.sql

## migrate
WRENCH_OPTION := --project ${SPANNER_PROJECT} --instance ${SPANNER_INSTANCE} --database ${SPANNER_DATABASE}

.PHONY: wrench/create
wrench/create: ## create spanner database
	$(GOBIN)/wrench create $(WRENCH_OPTION) --directory db/spanner

.PHONY: wrench/drop
wrench/drop: ## drop database in spanner
	$(GOBIN)/wrench drop $(WRENCH_OPTION) --directory db/spanner

.PHONY: wrench/reset
wrench/reset: ## drop the database and then re-create
	$(GOBIN)/wrench reset $(WRENCH_OPTION) --directory db/spanner

.PHONY: wrench/load-schema
wrench/load-schema:  ## update schema.sql from current database
	$(GOBIN)/wrench load $(WRENCH_OPTION) --directory db/spanner

.PHONY: wrench/apply
wrench/apply:  ## apply single DML
	$(GOBIN)/wrench apply $(WRENCH_OPTION) --dml $(DML)

.PHONY: wrench/migrate
wrench/migrate:  ## migrate by current migration files
	$(GOBIN)/wrench migrate up $(WRENCH_OPTION)

.PHONY: wrench/migrate-create
wrench/migrate-create:  ## create a spanner's migration file
	$(GOBIN)/wrench migrate create $(WRENCH_OPTION) --directory db/spanner

.PHONY: wrench/migrate-set
wrench/migrate-set:  ## update migration version and clear dirty flag, but don't run migration. argument: version=''
	$(GOBIN)/wrench migrate set $(WRENCH_OPTION) $(version)

.PHONY: wrench/migrate-version
wrench/migrate-version:  ## show current migration version
	$(GOBIN)/wrench migrate version $(WRENCH_OPTION)

## yo
YO_OPTION := ${SPANNER_PROJECT} ${SPANNER_INSTANCE} ${SPANNER_DATABASE} --ignore-tables SchemaMigrations
YO_DIR := ./internal/driver/spanner/repository
DOMAIN_DIR := ./internal/domain/repository

.PHONY: yo/generate
yo/generate: ## generate yo template
	@rm -f $(YO_DIR)/*.yo.go
	@rm -f $(DOMAIN_DIR)/*.yo.go
	@$(GOBIN)/yo $(YO_OPTION) --template-path $(YO_DIR)/templates/dto -o $(YO_DIR) --custom-types-file $(YO_DIR)/templates/custom_column_types.yml
	@$(GOBIN)/yo $(YO_OPTION) --template-path $(YO_DIR)/templates/repository -o $(YO_DIR) --custom-types-file $(YO_DIR)/templates/custom_column_types.yml --suffix _repository.yo.go
	@$(GOBIN)/yo $(YO_OPTION) --template-path $(YO_DIR)/templates/domain -o $(DOMAIN_DIR) --custom-types-file $(YO_DIR)/templates/custom_column_types.yml
