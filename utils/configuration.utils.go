package utils

import "os"

const DEVELOPMENT string = "DEVELOPMENT"

func GetEnvironmentVariable(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetEnvironment() string {
	return GetEnvironmentVariable("ENVIRONMENT", DEVELOPMENT)
}
