-- name: FollowFeed :one
INSERT INTO feed_user (user_id, feed_id, updated_at)
VALUES($1, $2, now())
RETURNING *;

-- name: DeleteFollow :exec
delete from feed_user where id = $1;

-- name: GetAllFollows :many
select * from feed_user where user_id = $1;