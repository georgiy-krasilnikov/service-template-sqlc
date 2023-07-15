package db

import (
	"os"

	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

const defaultMigrationsDir = "../migrations"
const migrationsDirEnvName = "MIGRATIONS_DIR"

func Migrate(connString string) error {
	var migrationsDir string
	if md := os.Getenv(migrationsDirEnvName); len(md) == 0 {
		migrationsDir = defaultMigrationsDir
	} else {
		migrationsDir = md
	}

	conn, err := goose.OpenDBWithDriver("postgres", connString)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		return err
	}

	if err := goose.Up(conn, migrationsDir); err != nil {
		return err
	}

	return nil
}
