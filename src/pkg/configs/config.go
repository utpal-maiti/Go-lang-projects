package configs

import (
	"log"
	"my-fiber-app/mod/pkg/utils"

	"github.com/joho/godotenv"
)

// Init loads environment variables and performs initial setup
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

// Getter functions for environment variables
func AppPort() string {
	return utils.GetEnv("APP_PORT")
}

func JWTSecret() string {
	return utils.GetEnv("JWT_SECRET")
}

func DatabaseURL() string {
	return utils.GetEnv("DATABASE_URL")
}

func RedisAddr() string {
	return utils.GetEnv("REDIS_ADDR")
}

func RedisPassword() string {
	return utils.GetEnv("REDIS_PASSWORD")
}
