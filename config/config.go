package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort string
	DBConn   string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return &Config{}, errors.New("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbConn := os.Getenv("PG_CONN")

	return &Config{
		HttpPort: port,
		DBConn:   dbConn,
	}, nil
}
