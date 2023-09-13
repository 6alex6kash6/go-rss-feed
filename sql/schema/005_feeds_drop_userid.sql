-- +goose Up
ALTER TABLE "feeds"
DROP COLUMN if exists "user_id";

-- +goose Down
ALTER TABLE "feeds"
ADD COLUMN user_id SERIAL NOT NULL references users (id) on delete cascade;