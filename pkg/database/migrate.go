package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

func Open(ctx context.Context, dsn string, migrationPath string) (*sql.DB, error) {
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
	if migrationPath != "" {
		ctx.L().Info("Running DB migration")
		// Setup DB migration
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return nil, err
		}
		m, err := migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s", migrationPath),
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

func RunTransaction(ctx context.Context, db *sql.DB, f func(context.Context, *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	var pErr, fErr error
	func() {
		defer func() {
			if r := recover(); r != nil {
				pErr = fmt.Errorf("Recovered tx: %#v", r)
			}
		}()
		fErr = f(ctx, tx)
	}()
	if pErr != nil {
		if err := tx.Rollback(); err != nil {
			ctx.L().Error("failed to rollback", zap.Error(err))
		}
		return pErr
	}
	if fErr != nil {
		if err := tx.Rollback(); err != nil {
			ctx.L().Error("failed to rollback", zap.Error(err))
		}
		return fErr
	}
	return tx.Commit()
}
