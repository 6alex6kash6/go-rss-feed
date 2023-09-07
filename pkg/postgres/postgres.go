package postgres

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"

	"github.com/6alex6kash6/go-rss-feed/internal/database"
)

type DB struct {
	Queries *database.Queries
}

func New(dbConn string) *DB {
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		slog.Error("Unable to connect to db")
	}
	return &DB{
		Queries: database.New(db),
	}
}
