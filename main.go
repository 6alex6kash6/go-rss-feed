package main

import (
	"log/slog"

	"github.com/6alex6kash6/go-rss-feed/config"
	"github.com/6alex6kash6/go-rss-feed/internal/app"
)

func main() {
	config, err := config.NewConfig()

	if err != nil {
		slog.Error("Config error: ", err)
	}
	app.Run(config)
}
