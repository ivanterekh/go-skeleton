package env

import (
	"os"
	"strings"
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
