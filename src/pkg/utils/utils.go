package utils

import (
	"log"
	"os"
)

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is not set, it logs a warning.
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Environment variable %s is not set", key)
	}
	return value
}
