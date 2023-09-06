package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	v1 "github.com/6alex6kash6/go-rss-feed/internal/controller/http/v1"
	"github.com/6alex6kash6/go-rss-feed/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbConn := os.Getenv("PG_CONN")
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		slog.Error("Unable to connect to db")
	}
	dbQueries := database.New(db)

	apiConfig := v1.NewApiConfig(dbQueries)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	api := chi.NewRouter()
	api.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		type res struct {
			Status string `json:"status"`
		}
		v1.RespondWithJSON(w, http.StatusOK, res{Status: "ok"})
	})
	api.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		v1.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	})

	api.Post("/user", apiConfig.CreateUser)

	r.Mount("/v1", api)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	slog.Info("Server started on port:", port)
	srv.ListenAndServe()
}
