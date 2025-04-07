package config

import (
	"gopkg.in/guregu/null.v4"
)

var Elasticsearch elasticsearch

type elasticsearch struct {
	Url      null.String
	Key      null.String
	Category null.String
}

func (i *elasticsearch) LoadFromEnvironment() {
	replaceWithEnvVariableString(&i.Url, "ELASTICSEARCH_URL")
	replaceWithEnvVariableString(&i.Key, "ELASTICSEARCH_KEY")
	replaceWithEnvVariableString(&i.Category, "ELASTICSEARCH_CATEGORY")
}
