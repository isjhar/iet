package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"github://mattes:personal-access-token@mattes/migrate_test",
		"postgres://localhost:5432/database?sslmode=enable")
	if err != nil {
		log.Panic(err)
	}
	err = m.Down()
	if err != nil {
		log.Panic(err)
	}
	err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		log.Panic(err)
	}
}
