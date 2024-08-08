package utils

import (
	"fmt"
	"os"

    "github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	port := os.Getenv("PORT")
    if port == "" {
        return nil, fmt.Errorf("PORT variable not setted in .env")
    }

    config := Config{
        PORT: port,
    }

	return &config, nil
}
