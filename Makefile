vet: generate
	buf lint
	go vet ./...

generate:
	sqlc generate
	buf generate

build-authentication:
	docker build -f build/authentication-service.dockerfile -t authentication:latest .

build-broker:
	docker build -f build/broker-service.dockerfile -t broker:latest .

build: build-authentication build-broker