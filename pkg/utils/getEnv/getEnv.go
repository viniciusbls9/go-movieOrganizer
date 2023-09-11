package getEnv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() string {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the env")
		return ""
	}

	return dbURL
}
