package db

import (
	"embed"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*
var migrations embed.FS

// Migrate performs PostgreSQL migrations from the migrations/ directory, using
// the provided connection string
func Migrate(dsn string) (int, error) {
	db, err := New(dsn)

	if err != nil {
		return 0, err
	}

	defer func() {
		if db.Sqlx == nil {
			return
		}

		if err := db.Sqlx.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrations,
		Root:       "migrations",
	}

	return migrate.Exec(db.Sqlx.DB, "postgres", migrations, migrate.Up)
}
