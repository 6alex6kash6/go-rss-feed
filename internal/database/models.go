// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"database/sql"
	"time"
)

type Feed struct {
	ID        int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Name      string
	Url       sql.NullString
}

type FeedUser struct {
	ID        int32
	UserID    int32
	FeedID    int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type User struct {
	ID        int32
	CreatedAt sql.NullTime
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
