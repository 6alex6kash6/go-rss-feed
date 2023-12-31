package v1

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/6alex6kash6/go-rss-feed/internal/database"
	"github.com/go-chi/chi/v5"
)

type userInput struct {
	Name string `json:"name"`
}

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	input := &userInput{}
	if err := ParseBody(r, input); err != nil {
		RespondWithError(w, 500, err.Error())
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		UpdatedAt: time.Now().UTC(),
		Name:      input.Name,
	})
	if err != nil {
		slog.Error("Error: ", err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	RespondWithJSON(w, http.StatusCreated, user)
}

func (cfg *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserKey)
	RespondWithJSON(w, http.StatusOK, user)

}

func (cfg *ApiConfig) NewUserRoutes(api *chi.Mux) {
	api.Post("/users", cfg.CreateUser)
	api.Get("/users", AuthMiddleware(cfg.GetUser, cfg.DB))
}
