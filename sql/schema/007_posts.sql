-- +goose Up
create table
    posts (
        id SERIAL PRIMARY KEY,
        created_at timestamptz (0) NULL DEFAULT now (),
        updated_at timestamptz (0) NULL,
        title TEXT NOT NULL,
        url TEXT NULL unique,
        description TEXT NULL,
        published_at timestamptz (0) NULL,
        feed_id SERIAL NOT NULL references feeds (id) on delete cascade
    );

-- +goose Down
DROP TABLE posts;