package v1

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/6alex6kash6/go-rss-feed/internal/database"
	"github.com/go-chi/chi/v5"
)

type feedInput struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (cfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserKey).(database.User)
	input := &feedInput{}
	if err := ParseBody(r, input); err != nil {
		RespondWithError(w, 500, err.Error())
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:   input.Name,
		Url:    sql.NullString{String: input.Url, Valid: true},
		UserID: user.ID,
	})
	if err != nil {
		slog.Error("Error: ", err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	RespondWithJSON(w, http.StatusCreated, feed)

}

func (cfg *ApiConfig) NewFeedRoutes(api *chi.Mux) {
	api.Post("/feeds", AuthMiddleware(cfg.CreateFeed, cfg.DB))
}
