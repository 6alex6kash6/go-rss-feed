package v1

import "github.com/6alex6kash6/go-rss-feed/internal/database"

type ApiConfig struct {
	DB *database.Queries
}

func NewApiConfig(db *database.Queries) *ApiConfig {
	return &ApiConfig{
		DB: db,
	}
}
