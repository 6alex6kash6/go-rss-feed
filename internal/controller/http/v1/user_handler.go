package v1

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/6alex6kash6/go-rss-feed/internal/database"
)

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := &parameters{}
	if err := ParseBody(r, params); err != nil {
		RespondWithError(w, 500, err.Error())
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		slog.Error("Error: ", err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	RespondWithJSON(w, http.StatusCreated, user)
}
