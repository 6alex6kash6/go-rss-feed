package app

import (
	"github.com/6alex6kash6/go-rss-feed/config"
	v1 "github.com/6alex6kash6/go-rss-feed/internal/controller/rest"
	"github.com/6alex6kash6/go-rss-feed/pkg/httpserver"
	"github.com/6alex6kash6/go-rss-feed/pkg/postgres"
)

func Run(cfg *config.Config) {
	db := postgres.New(cfg.DBConn)

	apiConfig := v1.NewApiConfig(db.Queries)
	router := v1.NewRouter(apiConfig)
	server := httpserver.New(router.Handler, cfg.HttpPort)
	server.Start()

}