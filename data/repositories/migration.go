package repositories

import (
	"isjhar/template/echo-golang/utils"

	"github.com/golang-migrate/migrate/v4"
)

const migrationPath = ""

func Migrate() error {
	m, err := migrate.New(
		utils.GetEnvironmentVariable("MIGRATION_PATH", migrationPath),
		GetDataSourceName())
	if err != nil {
		return err
	}
	err = m.Down()
	if err != nil {
		return err
	}
	err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		return err
	}
	return nil
}
