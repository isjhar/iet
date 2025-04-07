package config

import (
	"gopkg.in/guregu/null.v4"
)

var Swagger swagger

type swagger struct {
	Title  null.String `json:"title"`
	Scheme null.String `json:"scheme"`
}

func (i *swagger) LoadFromEnvironment() {
	replaceWithEnvVariableString(&i.Title, "SWAGGER_TITTLE")
	replaceWithEnvVariableString(&i.Scheme, "SCHEME")
}
