package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	envVarName = "ENV"
	prod       = "prod"
	dev        = "dev"
	staging    = "staging"
)

var env string

func init() {
	env = os.Getenv(envVarName)
	if env == "" {
		env = dev
	}
}

// IsProd returns true when app is running in production environment.
func IsProd() bool {
	return strings.ToLower(env) == prod
}

// IsDev returns true when app is running in development environment.
func IsDev() bool {
	return strings.ToLower(env) == dev
}

// IsStaging returns true when app is running in staging environment.
func IsStaging() bool {
	return strings.ToLower(env) == staging
}

// GetInt returns environment variable converted to int if it
// exists and can be parsed or the default value in other case.
func GetInt(name string, defaultVal int) int {
	valStr, ok := os.LookupEnv(name)
	if !ok {
		return defaultVal
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetFloat64 returns environment variable converted to float64 if it
// exists and can be parsed or the default value in other case.
func GetFloat64(name string, defaultVal float64) float64 {
	valStr, ok := os.LookupEnv(name)
	if !ok {
		return defaultVal
	}

	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetBool returns environment variable converted to bool if it
// exists and can be parsed or the default value in other case.
func GetBool(name string, defaultVal bool) bool {
	valStr, ok := os.LookupEnv(name)
	if !ok {
		return defaultVal
	}

	val, err := strconv.ParseBool(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetString returns environment variable if it exists
// or the default value in other case.
func GetString(name string, defaultVal string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultVal
	}

	return val
}

// GetDuration returns environment variable if it exists
// or the default value in other case.
func GetDuration(name string, defaultVal time.Duration) time.Duration {
	valStr, ok := os.LookupEnv(name)
	if !ok {
		return defaultVal
	}

	val, err := time.ParseDuration(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}
