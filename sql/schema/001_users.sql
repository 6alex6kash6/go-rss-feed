-- +goose Up
CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        created_at timestamptz (0) NULL DEFAULT now (),
        updated_at TIMESTAMP NOT NULL,
        name TEXT NOT NULL
    );

-- +goose Down
DROP TABLE users;