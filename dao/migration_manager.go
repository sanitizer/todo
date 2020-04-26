package dao

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"log"
)

func MigrateSchema() {
	db := Connection()
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	checkForErrorAndFail(err)

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	checkForErrorAndFail(err)

	m.Steps(2)
}

func checkForErrorAndFail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
