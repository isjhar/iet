package config

import (
	"encoding/json"
	"os"
)

type config struct {
	Database      database
	Migration     migration
	Swagger       swagger
	Elasticsearch elasticsearch
}

func LoadConfig() {
	cfg, err := loadFromJson()
	if err == nil {
		Database = cfg.Database
		Migration = cfg.Migration
		Swagger = cfg.Swagger
		Elasticsearch = cfg.Elasticsearch
	}

	Swagger.LoadFromEnvironment()
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
