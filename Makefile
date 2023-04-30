vet: buf generate
	$(BUF) lint
	go vet ./...

web/node_module: web/package.json
	cd web;npm install --silent

generate: sqlc buf web/node_module
	$(SQLC) generate
	$(BUF) generate
	sed -i '' 's/.js\"/\"/g' web/src/client/*/*/*.ts

build-web: web/node_module
	docker build -f build/Dockerfile -t igru:latest .

build: build-web

test-env:
	@echo "Creating Mock Database..."
	docker-compose -f "./e2e/docker-compose.yml" up -d

test-services:
	@echo "Running Tests"
	go test -v ./e2e/gardens/services

test-run: test-services

test-clean:
	@echo "Cleaning Up Testing Environment..."
	docker-compose -f "./e2e/docker-compose.yml" down
	@echo "Clearing Database..."
	rm -rf ./e2e/db-data

test-auto: test-env test-run test-clean

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
SQLC ?= $(LOCALBIN)/sqlc
SQLC_VERSION ?= v1.18.0
BUF ?= $(LOCALBIN)/buf
BUF_VERSION ?= v1.17.0

.PHONY: sqlc
sqlc: $(SQLC)
$(SQLC): $(LOCALBIN)
	test -s $(SQLC)|| GOBIN=$(LOCALBIN) go install github.com/kyleconroy/sqlc/cmd/sqlc@$(SQLC_VERSION)

.PHONY: buf
buf: $(BUF)
$(BUF): $(LOCALBIN)
	test -s $(BUF)|| GOBIN=$(LOCALBIN) go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)