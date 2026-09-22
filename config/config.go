package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVal struct{
	REDIS_URL string
}

func godotenvvariable(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Env was not loaded")
	}

	return os.Getenv(key);
}

func GetEnv(key string) string {
	return godotenvvariable(key);
}