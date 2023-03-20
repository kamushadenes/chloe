package memory

import (
	"context"
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

var db *sql.DB

var DefaultMigrationsTable = "chloe_migrations"

//go:embed migrations/*.sql
var migrations embed.FS

func Setup(ctx context.Context) error {
	var err error
	db, err = sql.Open("sqlite3", "chloe.db")
	if err != nil {
		return err
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{MigrationsTable: DefaultMigrationsTable})
	if err != nil {
		return err
	}

	d, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", d, "chloe", driver)
	if err != nil {
		return err
	}
	m.Up()

	return nil
}

func Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			db.Close()
			return
		}
	}
}
