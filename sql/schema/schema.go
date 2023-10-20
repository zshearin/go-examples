package schema

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// following is the way to embed files into go code - needs to be no space between comment and go (//go:embed, not // go:embed)

//go:embed migrations/*.sql
var fs embed.FS
var path = "migrations"

// Version - active schema version (corresponds to sql statements in migrations folder)
var Version uint = 2

func NewMigrator(dsn string) (*migrate.Migrate, error) {
	return NewEmbeddedMigrator(dsn, fs, path)
}

func NewEmbeddedMigrator(dsn string, fs embed.FS, path string) (*migrate.Migrate, error) {
	driver, err := iofs.New(fs, path)
	if err != nil {
		return nil, err
	}
	return migrate.NewWithSourceInstance("iofs", driver, dsn)
}
