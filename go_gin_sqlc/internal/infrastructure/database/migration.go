package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	migrate "github.com/rubenv/sql-migrate"
)

func execMigrations(db *sql.DB, direction migrate.MigrationDirection) error {
	migrate.SetTable("migrations")

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	migrationCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	migrationCount, err := migrate.ExecContext(migrationCtx, db, "postgres", migrations, direction)
	if err != nil {
		return err
	}

	log.Printf("Applied %d migrations\n", migrationCount)

	return nil
}

func MigrateUp(db *sql.DB) error {
	return execMigrations(db, migrate.Up)
}

func MigrateDown(db *sql.DB) error {
	return execMigrations(db, migrate.Down)
}
