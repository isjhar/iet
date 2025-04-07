package config

import (
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
	replaceWithEnvVariableString(&i.Host, "DB_HOST")
	replaceWithEnvVariableString(&i.Port, "DB_PORT")
	replaceWithEnvVariableString(&i.User, "DB_USER")
	replaceWithEnvVariableString(&i.Password, "DB_PASSWORD")
	replaceWithEnvVariableString(&i.Database, "DB_NAME")
}
