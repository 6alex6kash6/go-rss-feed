// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO
    posts (title, url, description, published_at, feed_id, updated_at) 
    VALUES($1, $2, $3, $4, $5, now())
    RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	Title       string
	Url         sql.NullString
	Description sql.NullString
	PublishedAt sql.NullTime
	FeedID      int32
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getUserPosts = `-- name: GetUserPosts :many
SELECT posts.id, posts.created_at, posts.updated_at, posts.title, posts.url, posts.description, posts.published_at, posts.feed_id FROM posts
JOIN feed_user ON feed_user.feed_id = posts.feed_id
WHERE feed_user.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2
`

type GetUserPostsParams struct {
	UserID int32
	Limit  int32
}

func (q *Queries) GetUserPosts(ctx context.Context, arg GetUserPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getUserPosts, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
