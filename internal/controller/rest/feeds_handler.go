package v1

import (
	"database/sql"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/6alex6kash6/go-rss-feed/internal/database"
	"github.com/go-chi/chi/v5"
)

type feedInput struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type followFeedInput struct {
	FeedId int32 `json:"feed_id"`
}

type createFeedResponse struct {
	Feed       database.Feed     `json:"feed"`
	FeedFollow database.FeedUser `json:"feed_follow"`
}

func (cfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserKey).(database.User)
	input := &feedInput{}
	if err := ParseBody(r, input); err != nil {
		RespondWithError(w, 500, err.Error())
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name: input.Name,
		Url:  sql.NullString{String: input.Url, Valid: true},
	})
	if err != nil {
		slog.Error("Error: ", err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	follow, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		slog.Error("Error: ", err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't follow feed")
		return
	}
	RespondWithJSON(w, http.StatusCreated, createFeedResponse{
		Feed:       feed,
		FeedFollow: follow,
	})

}

func (cfg *ApiConfig) FollowFeed(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserKey).(database.User)
	input := &followFeedInput{}
	if err := ParseBody(r, input); err != nil {
		RespondWithError(w, 500, err.Error())
		return
	}
	follow, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		UserID: user.ID,
		FeedID: input.FeedId,
	})
	if err != nil {
		slog.Error("Error: ", err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't follow feed")
		return
	}
	RespondWithJSON(w, http.StatusCreated, follow)
}

func (cfg *ApiConfig) DeleteFollowFeed(w http.ResponseWriter, r *http.Request) {
	feedFollowId, err := strconv.Atoi(chi.URLParam(r, "feedFollowId"))
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Wrong feedFollowId")
	}

	if err := cfg.DB.DeleteFollow(r.Context(), int32(feedFollowId)); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable delete follow")
	}
}

func (cfg *ApiConfig) GetFollowFeeds(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserKey).(database.User)

	follows, err := cfg.DB.GetAllFollows(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to get follows")
	}
	RespondWithJSON(w, http.StatusOK, follows)
}

func (cfg *ApiConfig) GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		slog.Error("Error: ", err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}
	RespondWithJSON(w, http.StatusOK, feeds)
}

func (cfg *ApiConfig) NewFeedRoutes(api *chi.Mux) {
	api.Post("/feeds", AuthMiddleware(cfg.CreateFeed, cfg.DB))
	api.Get("/feeds/all", cfg.GetAllFeeds)
	api.Post("/feeds/follows", AuthMiddleware(cfg.FollowFeed, cfg.DB))
	api.Delete("/feeds/follows/{feedFollowId}", AuthMiddleware(cfg.DeleteFollowFeed, cfg.DB))
	api.Get("/feeds/follows", AuthMiddleware(cfg.GetFollowFeeds, cfg.DB))
}
