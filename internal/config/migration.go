package config

import (
	"github.com/isjhar/iet/pkg"
	"gopkg.in/guregu/null.v4"
)

var Migration migration

type migration struct {
	Path null.String
}

func (i *migration) LoadFromEnvironment() {
	i.Path = null.StringFrom(pkg.GetEnvironmentVariable("PACKAGE_PATH", ""))
}
