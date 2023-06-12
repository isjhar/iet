package repositories

import (
	"errors"
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
	version, _, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return err
	}
	if version > 0 {
		err = m.Down()
		if err != nil {
			return err
		}
	}

	err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		return err
	}
	return nil
}
