package repositories

import (
	"errors"
	"fmt"
	"isjhar/template/echo-golang/utils"

	"github.com/golang-migrate/migrate/v4"
)

const migrationPath = ""

func MigrateDatabase() error {
	return Migrate(getMigrationPath())
}

func MigrateSeed() error {
	return Migrate(getSeedPath())
}

func Migrate(migrationPath string) error {
	m, err := migrate.New(
		migrationPath,
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

func getMigrationPath() string {
	packagePath := utils.GetEnvironmentVariable("PACKAGE_PATH", packagePath)
	return fmt.Sprintf("file://%s/data/repositories/migrations", packagePath)
}

func getSeedPath() string {
	packagePath := utils.GetEnvironmentVariable("PACKAGE_PATH", packagePath)
	return fmt.Sprintf("file://%s/data/repositories/seeds", packagePath)
}
