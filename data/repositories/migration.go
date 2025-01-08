package repositories

import (
	"errors"
	"fmt"
	"isjhar/template/echo-golang/domain/entities"
	"isjhar/template/echo-golang/utils"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var isRefreshed = false

func Refresh() error {
	if isRefreshed {
		return nil
	}
	return refresh()
}

func ForceRefresh() error {
	return refresh()
}

func refresh() error {
	databaseMigrate, err := CreateMigrate(getMigrationPath())
	if err != nil {
		return err
	}

	version, _, err := databaseMigrate.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return err
	}
	if version > 0 {
		err = databaseMigrate.Down()
		if err != nil {
			return err
		}
	}

	err = databaseMigrate.Up()
	if err != nil {
		return err
	}
	isRefreshed = true
	return nil
}

func MigrateDatabase() error {
	m, err := CreateMigrate(getMigrationPath())
	if err != nil {
		return err
	}
	migrationVersion, err := getLastDatabaseVersion()
	if err != nil {
		return err
	}
	err = m.Migrate(migrationVersion) // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		return err
	}
	return nil
}

func getLastDatabaseVersion() (uint, error) {
	var result uint = 1
	files, err := os.ReadDir(getMigrationFolderPath())
	if err != nil {
		return result, err
	}
	versionRegex, err := regexp.Compile("^99999[0-9]+.*")
	if err != nil {
		return result, err
	}
	for _, file := range files {
		filename := file.Name()
		if !versionRegex.MatchString(filename) {
			versionParts := strings.Split(filename, "_")
			if len(versionParts) == 0 {
				return result, entities.InternalServerError
			}

			u64, err := strconv.ParseUint(versionParts[0], 10, 32)
			if err != nil {
				continue
			}
			if u64 > uint64(result) {
				result = uint(u64)
			}
		}
	}

	return result, nil
}

func MigrateSeed() error {
	m, err := CreateMigrate(getMigrationPath())
	if err != nil {
		return err
	}
	err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		return err
	}
	return nil
}

func CreateMigrate(migrationPath string) (*migrate.Migrate, error) {
	dataSourceName := GetDataSourceName()
	return migrate.New(
		migrationPath,
		dataSourceName)

}

func getMigrationPath() string {
	return fmt.Sprintf("file://%s", getMigrationFolderPath())
}

func getMigrationFolderPath() string {
	packagePath := utils.GetEnvironmentVariable("PACKAGE_PATH", packagePath)
	return fmt.Sprintf("%s/data/repositories/migrations", packagePath)
}
