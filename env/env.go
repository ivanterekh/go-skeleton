package env

import "os"

var env = os.Getenv("ENV")

// Is tells if app is running in mode environment.
func Is(mode string) bool {
	return env == mode
}
