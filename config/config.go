package config

import (
	"errors"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort    string
	DBConn      string
	RunScrapper bool
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return &Config{}, errors.New("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbConn := os.Getenv("PG_CONN")
	runScrapper, err := strconv.ParseBool(os.Getenv("RUN_SCRAPPER"))
	if err != nil {
		slog.Error("Cannot parse RUN_SCRAPPER")
	}
	return &Config{
		HttpPort:    port,
		DBConn:      dbConn,
		RunScrapper: runScrapper,
	}, nil
}
