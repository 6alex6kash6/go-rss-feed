-- name: CreateFeed :one
INSERT INTO
    feeds (name, url, user_id, updated_at) 
    VALUES($1, $2, $3, now())
    RETURNING *;