package config

import (
	"github.com/isjhar/iet/pkg"
	"gopkg.in/guregu/null.v4"
)

var Database database

type database struct {
	Host     null.String
	Port     null.String
	User     null.String
	Password null.String
	Database null.String
}

func (i *database) LoadFromEnvironment() {
	i.Host = null.StringFrom(pkg.GetEnvironmentVariable("DB_HOST", ""))
	i.Port = null.StringFrom(pkg.GetEnvironmentVariable("DB_PORT", ""))
	i.User = null.StringFrom(pkg.GetEnvironmentVariable("DB_USER", ""))
	i.Password = null.StringFrom(pkg.GetEnvironmentVariable("DB_PASSWORD", ""))
	i.Database = null.StringFrom(pkg.GetEnvironmentVariable("DB_NAME", ""))
}
