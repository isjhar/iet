package config

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/isjhar/iet/pkg"
	"gopkg.in/guregu/null.v4"
)

type config struct {
	Database      database
	Migration     migration
	Swagger       swagger
	Elasticsearch elasticsearch
	Jwt           jwt
}

func LoadConfig() {
	cfg, err := loadFromJson()
	if err == nil {
		Database = cfg.Database
		Migration = cfg.Migration
		Swagger = cfg.Swagger
		Elasticsearch = cfg.Elasticsearch
		Jwt = cfg.Jwt
	}

	Database.LoadFromEnvironment()
	Migration.LoadFromEnvironment()
	Swagger.LoadFromEnvironment()
	Elasticsearch.LoadFromEnvironment()
	Jwt.LoadFromEnvironment()
}

func loadFromJson() (config, error) {
	var cfg config

	file, err := os.Open("../../config.json")
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func replaceWithEnvVariableString(configValue *null.String, envKey string) {
	title := pkg.GetEnvironmentVariable(envKey, "")
	if title != "" {
		*configValue = null.StringFrom(title)
	}
}

func replaceWithEnvVariableInt(configValue *null.Int, envKey string) {
	stringVal := pkg.GetEnvironmentVariable(envKey, "")
	if stringVal != "" {
		intVal, _ := strconv.ParseInt(stringVal, 10, 64)
		*configValue = null.IntFrom(intVal)
	}
}
