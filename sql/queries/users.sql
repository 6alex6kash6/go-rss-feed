-- name: CreateUser :one
INSERT INTO users (updated_at, name)
VALUES ($1, $2)
RETURNING *;