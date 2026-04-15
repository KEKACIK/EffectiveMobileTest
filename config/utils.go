package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func initConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Fatal error in environment variables")
	}
}

func getStrEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getIntEnv(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, _ := strconv.Atoi(value)
		return intValue
	}

	return defaultVal
}

func getBoolEnv(key string, defaultVal bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "1"
	}

	return defaultVal
}
