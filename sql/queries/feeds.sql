-- name: CreateFeed :one
INSERT INTO
    feeds (name, url, updated_at) 
    VALUES($1, $2, now())
    RETURNING *;

-- name: GetAllFeeds :many
select * from feeds;

-- name: UpdateLastFetched :one
update feeds set last_fetched_at = now() where id = $1 RETURNING *;

-- name: GetFeedsToFetch :many
select * from feeds order by last_fetched_at asc nulls first limit $1;