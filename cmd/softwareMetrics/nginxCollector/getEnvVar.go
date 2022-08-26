package nginxCollector

import (
	"log"
	"os"
	"strconv"
	"time"
)

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func getEnvUint(key string, defaultValue uint) uint {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Fatalf("Environment variable value for %s must be an uint: %v", key, err)
	}
	return uint(i)
}

func getEnvBool(key string, defaultValue bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatalf("Environment variable value for %s must be a boolean: %v", key, err)
	}
	return b
}

func getEnvPositiveDuration(key string, defaultValue time.Duration) positiveDuration {
	value, ok := os.LookupEnv(key)
	if !ok {
		return positiveDuration{defaultValue}
	}

	posDur, err := parsePositiveDuration(value)
	if err != nil {
		log.Fatalf("Environment variable value for %s must be a positive duration: %v", key, err)
	}
	return posDur
}

func getEnvConstLabels(key string, defaultValue map[string]string) constLabel {
	value, ok := os.LookupEnv(key)
	if !ok {
		return constLabel{defaultValue}
	}

	cLabel, err := parseConstLabels(value)
	if err != nil {
		log.Fatalf("Environment variable value for %s must be a const label or a list of const labels: %v", key, err)
	}
	return cLabel
}
