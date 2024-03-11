package config

import "os"

// IsProduction returns true if APP_ENV is production
func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}
