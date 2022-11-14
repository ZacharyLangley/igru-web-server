package database

import (
	"database/sql"
	"errors"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open(ctx context.Context, dsn string, disableMigration bool) (*sql.DB, error) {
	// Connect to DB
	ctx.L().Debug("Opening connection to DB")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// Check DB connection
	ctx.L().Debug("Testing connection to DB")
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	if !disableMigration {
		ctx.L().Info("Running DB migration")
		// Setup DB migration
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return nil, err
		}
		m, err := migrate.NewWithDatabaseInstance(
			"file:///migrations",
			"pgx", driver)
		if err != nil {
			return nil, err
		}
		// Run DB migration
		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	} else {
		ctx.L().Info("Skipping DB migration")
	}
	return db, nil
}
