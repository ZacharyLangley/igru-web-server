vet: generate
	buf lint
	go vet ./...

generate:
	sqlc generate
	buf generate
