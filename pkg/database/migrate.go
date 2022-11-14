package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open(ctx context.Context, dsn string, disableMigration bool) (*sql.DB, error) {
	// Connect to DB
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// Check DB connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	if !disableMigration {
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
	}
	return db, nil
}
