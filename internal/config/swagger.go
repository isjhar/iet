package config

import (
	"github.com/isjhar/iet/pkg"
	"gopkg.in/guregu/null.v4"
)

var Swagger swagger

type swagger struct {
	Title  null.String `json:"title"`
	Scheme null.String `json:"scheme"`
}

func (i *swagger) LoadFromEnvironment() {
	i.Title = null.StringFrom(pkg.GetEnvironmentVariable("SWAGGER_TITLE", ""))
	i.Scheme = null.StringFrom(pkg.GetEnvironmentVariable("SCHEME", ""))
}
