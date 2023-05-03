package repositories

import (
	"github.com/golang-migrate/migrate/v4"
)

func Migrate() error {
	m, err := migrate.New(
		"github://mattes:personal-access-token@mattes/migrate_test",
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
