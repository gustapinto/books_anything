package repository

import (
	"context"
	"database/sql"
)

type Migrator interface {
	Migrate() string
}

func AutoMigrate(db *sql.DB, migrators ...Migrator) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, migrator := range migrators {
		if _, err := tx.ExecContext(ctx, migrator.Migrate()); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
