package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	Port       int
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	return &Config{
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		Port:       8080,
	}, nil
}
