-- +goose Up
ALTER TABLE feeds
ADD COLUMN "last_fetched_at" timestamptz (0) NULL;

-- +goose Down
ALTER TABLE feeds
DROP COLUMN "last_fetched_at";