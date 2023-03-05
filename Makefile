vet: generate
	buf lint
	go vet ./...

web/node_module: web/package.json
	cd web;npm install

generate: web/node_module
	sqlc generate
	buf generate
	sed -i '' 's/.js\"/\"/g' web/src/client/*/*/*.ts

build-authentication:
	docker build -f build/Dockerfile -t authentication:latest .

build-garden:
	docker build -f build/Dockerfile -t garden:latest .

build-broker:
	docker build -f build/Ddockerfile -t broker:latest .

build-web: web/node_module
	BUILD_PATH=../cmd/ingress/public npm run build

build: build-authentication build-garden build-broker build-web

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