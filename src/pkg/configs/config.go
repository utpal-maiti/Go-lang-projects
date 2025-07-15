package configs

import (
	"log"
	"my-fiber-app/mod/pkg/utils"

	"github.com/joho/godotenv"
)

// Init loads environment variables and performs initial setup.
func Init() {
	// Load environment variables from .env file
	// This is useful for local development
	// In production, you might want to set these variables directly in the environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Configuration initialized")
}

// AppPort returns the application port from environment variables.
func AppPort() string {
	return utils.GetEnv("APP_PORT")
}

// JWTSecret returns the JWT secret from environment variables.
func JWTSecret() string {
	return utils.GetEnv("JWT_SECRET")
}

// DatabaseURL returns the database URL from environment variables.
func DatabaseURL() string {
	return utils.GetEnv("DATABASE_URL")
}

// RedisAddr returns the Redis address from environment variables.
func RedisAddr() string {
	return utils.GetEnv("REDIS_ADDR")
}

// RedisPassword returns the Redis password from environment variables.
func RedisPassword() string {
	return utils.GetEnv("REDIS_PASSWORD")
}
