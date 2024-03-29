# igru-web-server

## Commands

- Generate Go Code from ProtoFiles = "buf generate"
- Generate Models from ".sql" files = "sqlc generate"
- Generate BOTH Go Code and Models = "make generate"
- Generate Builds for all services = "make build"

## Local Web Development

1. Bring up docker-compose. `docker-compose up -d`
2. Tear down the ingress, it's only slowing you down. `docker-compose down ingress`
3. Start the ingress locally through the VSCode debugger.
4. Run local web development environment. `cd web && npm run`

## Utility Functions

### Auth

```sh
go run ./cmd/utils/add-user

go run ./cmd/utils/whoami $(go run ./cmd/utils/login)

```

### Gardens

**Be sure to add any necessary information such as Garden IDs before executing GET, UPDATE, DELETE operations.**

```sh
go run ./cmd/utils/gardens/createGarden

go run ./cmd/utils/gardens/updateGarden

go run ./cmd/utils/gardens/deleteGarden

go run ./cmd/utils/gardens/getGarden

go run ./cmd/utils/gardens/getGardens

```