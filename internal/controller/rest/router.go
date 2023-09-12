package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Router struct {
	Handler *chi.Mux
}

func NewRouter(apiConfig *ApiConfig) *Router {
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
	api.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		type res struct {
			Status string `json:"status"`
		}
		RespondWithJSON(w, http.StatusOK, res{Status: "ok"})
	})
	api.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	})
	{
		apiConfig.NewUserRoutes(api)
		apiConfig.NewFeedRoutes(api)
	}
	r.Mount("/v1", api)
	return &Router{
		Handler: r,
	}
}
