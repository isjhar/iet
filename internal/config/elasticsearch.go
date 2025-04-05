package config

import (
	"github.com/isjhar/iet/pkg"
	"gopkg.in/guregu/null.v4"
)

var Elasticsearch elasticsearch

type elasticsearch struct {
	Url      null.String
	Key      null.String
	Category null.String
}

func (i *elasticsearch) LoadFromEnvironment() {
	i.Url = null.StringFrom(pkg.GetEnvironmentVariable("ELASTICSEARCH_URL", ""))
	i.Key = null.StringFrom(pkg.GetEnvironmentVariable("ELASTICSEARCH_KEY", ""))
	i.Category = null.StringFrom(pkg.GetEnvironmentVariable("ELASTICSEARCH_CATEGORY", ""))
}
