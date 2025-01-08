package utils

import "os"

const DEVELOPMENT string = "develop"
const STAGING string = "staging"
const PROD string = "prod"
const LOCAL string = "LOCAL"

func GetEnvironmentVariable(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetEnvironment() string {
	return GetEnvironmentVariable("ENVIRONMENT", LOCAL)
}

func GetBaseUrl() string {
	return GetEnvironmentVariable("BASE_URL", "localhost:1323")
}
