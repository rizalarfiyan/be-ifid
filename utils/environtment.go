package utils

import (
	"os"
	"strconv"
	"time"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsTimeDuration(name string, defaultVal time.Duration) time.Duration {
	valStr := GetEnv(name, "")
	val, err := time.ParseDuration(valStr)
	if err == nil {
		return val
	}

	return defaultVal
}
