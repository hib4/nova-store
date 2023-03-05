package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed loading .env")
	}

	return os.Getenv(key)
}
