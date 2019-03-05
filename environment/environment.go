package environment

import "os"

func Get() string {
	env := os.Getenv("ENV")
	if env == "" {
		return "dev"
	}

	return env
}
