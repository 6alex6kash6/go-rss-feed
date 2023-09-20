package v1

import (
	"net/http"
	"strconv"

	"github.com/6alex6kash6/go-rss-feed/internal/database"
	"github.com/go-chi/chi/v5"
)

func (cfg *ApiConfig) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}
	user := r.Context().Value(UserKey).(database.User)

	posts, err := cfg.DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't get posts for user")
		return
	}
	RespondWithJSON(w, http.StatusOK, posts)
}

func (cfg *ApiConfig) NewPostsRoutes(api *chi.Mux) {
	api.Get("/posts", AuthMiddleware(cfg.GetUserPosts, cfg.DB))
}
