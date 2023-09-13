-- name: CreateFeed :one
INSERT INTO
    feeds (name, url, updated_at) 
    VALUES($1, $2, now())
    RETURNING *;

-- name: GetAllFeeds :many
select * from feeds;