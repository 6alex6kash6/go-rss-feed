-- +goose Up
create table
    feeds (
        id SERIAL PRIMARY KEY,
        created_at timestamptz (0) NULL DEFAULT now (),
        updated_at timestamptz (0) NULL,
        name TEXT NOT NULL,
        url TEXT unique
    );

-- +goose Down
DROP TABLE feeds