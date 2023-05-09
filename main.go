package main

import (
	"github.com/ZacharyLangley/igru-web-server/cmd"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	cmd.Execute()
}
