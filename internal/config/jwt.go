package config

import "gopkg.in/guregu/null.v4"

var Jwt jwt

type jwt struct {
	Secret               null.String
	AccessTokenExpiresIn null.Int
}

func (i *jwt) LoadFromEnvironment() {
	replaceWithEnvVariableString(&i.Secret, "JWT_SECRET")
	replaceWithEnvVariableInt(&i.AccessTokenExpiresIn, "JWT_ACCESS_TOKEN_EXPIRES_IN")
}
