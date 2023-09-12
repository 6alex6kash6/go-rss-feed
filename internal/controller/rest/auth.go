package v1

import (
	"context"
	"net/http"

	"log/slog"

	"github.com/6alex6kash6/go-rss-feed/internal/database"
)

type User string

const UserKey User = "user"

type authedHandler func(http.ResponseWriter, *http.Request)

func AuthMiddleware(next authedHandler, db *database.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		user, err := db.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			slog.Error("Error: ", err)
			RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
			return
		}
		ctx := context.WithValue(r.Context(), UserKey, user)
		next(w, r.WithContext(ctx))
	})
}
