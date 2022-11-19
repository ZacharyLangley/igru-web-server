package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

type Pool struct {
	*pgxpool.Pool
}

func Open(ctx context.Context, cfg config.Database) (*Pool, error) {
	dsn, err := cfg.DSN()
	if err != nil {
		return nil, err
	}
	// Connect to DB Pool
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
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
	if cfg.MigrationPath != "" {
		ctx.L().Info("Running DB migration")
		// Setup DB migration
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return nil, err
		}
		m, err := migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s", cfg.MigrationPath),
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

	return &Pool{pool}, nil
}

func (p *Pool) RunTransaction(ctx context.Context, f func(context.Context, pgx.Tx) error) error {
	// Acquire DB connection from pool
	db, err := p.Acquire(ctx)
	if err != nil {
		return err
	}
	defer db.Release()
	// Start transaction
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	// Run transaction function
	var pErr, fErr error
	func() {
		defer func() {
			if r := recover(); r != nil {
				pErr = fmt.Errorf("Recovered tx: %#v", r)
			}
		}()
		fErr = f(ctx, tx)
	}()
	// Rollback if required
	if pErr != nil {
		if err := tx.Rollback(ctx); err != nil {
			ctx.L().Error("failed to rollback", zap.Error(err))
		}
		return pErr
	}
	if fErr != nil {
		if err := tx.Rollback(ctx); err != nil {
			ctx.L().Error("failed to rollback", zap.Error(err))
		}
		return fErr
	}
	// Commit if no errors are found
	return tx.Commit(ctx)
}
