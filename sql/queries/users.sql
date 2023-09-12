-- name: CreateUser :one
INSERT INTO users (updated_at, name, api_key)
VALUES ($1, $2, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByApiKey :one
select * from users where api_key = $1;