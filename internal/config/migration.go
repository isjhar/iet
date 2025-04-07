package config

import (
	"gopkg.in/guregu/null.v4"
)

var Migration migration

type migration struct {
	Path null.String
}

func (i *migration) LoadFromEnvironment() {
	replaceWithEnvVariableString(&i.Path, "PACKAGE_PATH")
}
