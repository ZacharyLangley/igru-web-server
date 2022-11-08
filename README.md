# igru-web-server

`.sql` files are used to generate models and to run migrations

The services use those models to fulfil requests

```sh
go run ./cmd/utils/add-user
```

```sh
go run ./cmd/utils/whoami $(go run ./cmd/utils/login)
```
