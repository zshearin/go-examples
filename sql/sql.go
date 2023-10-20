package sql

import (
	"log"

	"github.com/zshearin/go-examples/sql/schema"
)

func DoExample() {

	// creating necessary mysql to interact with:
	// brew install mysql
	// brew services start
	// CREATE DATABASE my_database;

	// dsn := "username:password@tcp(ipAddressOfMySQL:3306)/database_name"
	dsn := "root@tcp(127.0.0.1:3306)/my_database"

	createAndRunSchemaMigrator(dsn)

}

func createAndRunSchemaMigrator(dsn string) {

	sqlAddress := "mysql://" + dsn + "?multiStatements=true"
	migrator, err := schema.NewMigrator(sqlAddress)
	if err != nil {
		log.Fatalf("failed creating sql migrator: %v\n", err)
	}

	migrator.Migrate(schema.Version)
}
