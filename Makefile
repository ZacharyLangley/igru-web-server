vet: generate
	buf lint
	go vet ./...

generate:
	sqlc generate
	buf generate

keys:
	mkdir -p .build/keys
	go run . auth create-key --file build/authentication/keys/id_ed25519

build-authentication:
	docker build -f build/authentication-service.dockerfile -t authentication:latest .

build-garden:
	docker build -f build/garden-service.dockerfile -t garden:latest .

build-broker:
	docker build -f build/broker-service.dockerfile -t broker:latest .

build: build-authentication build-garden build-broker

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