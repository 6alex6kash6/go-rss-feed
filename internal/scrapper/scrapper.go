package scrapper

import (
	"context"
	"fmt"
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

	slog.Info(item.Channel.Title)
}
