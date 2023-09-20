package scrapper

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"log/slog"

	"github.com/6alex6kash6/go-rss-feed/internal/database"
)

type Scrapper struct {
	DB *database.Queries
}

func NewScrapper(db *database.Queries) *Scrapper {
	return &Scrapper{
		DB: db,
	}
}

func (sc *Scrapper) Run(n int32) {
	var wg sync.WaitGroup
	slog.Info("Scrapper started for:", n)

	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		feeds, err := sc.DB.GetFeedsToFetch(context.Background(), n)
		if err != nil {
			fmt.Print(err)
		}
		for _, feed := range feeds {
			wg.Add(1)
			go sc.processFeed(feed, &wg)
		}
	}

	wg.Wait()
}

func (sc *Scrapper) processFeed(feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()
	slog.Info("Processing feed:", feed.ID)
	_, err := sc.DB.UpdateLastFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Print(err)
	}

	item, err := FetchFeed(feed.Url.String)
	if err != nil {
		slog.Error(err.Error())
	}

	for _, p := range item.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, p.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := sc.DB.CreatePost(context.Background(), database.CreatePostParams{
			Title: p.Title,
			Url: sql.NullString{
				String: p.Link,
				Valid:  true,
			},
			Description: sql.NullString{
				String: p.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			slog.Info("Couldn't create post: %v", err)
			continue
		}
	}
}
