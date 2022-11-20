# igru-web-server

## Commands:
- Generate Go Code from ProtoFiles = "buf generate"
- Generate Models from ".sql" files = "sqlc generate"

## Utility Functions

### Auth

```sh
go run ./cmd/utils/add-user
```

```sh
go run ./cmd/utils/whoami $(go run ./cmd/utils/login)
```

### Gardens

```sh
go run ./cmd/utils/gardens/createGarden

go run ./cmd/utils/gardens/updateGarden

go run ./cmd/utils/gardens/deleteGarden

go run ./cmd/utils/gardens/getGarden

go run ./cmd/utils/gardens/getGardens

```