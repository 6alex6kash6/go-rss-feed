-- name: CreatePost :one
INSERT INTO
    posts (title, url, description, published_at, feed_id, updated_at) 
    VALUES($1, $2, $3, $4, $5, now())
    RETURNING *;

-- name: GetUserPosts :many
SELECT posts.* FROM posts
JOIN feed_user ON feed_user.feed_id = posts.feed_id
WHERE feed_user.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;