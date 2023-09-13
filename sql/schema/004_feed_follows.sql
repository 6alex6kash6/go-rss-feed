-- +goose Up
create table
    feed_user (
        id SERIAL PRIMARY KEY,
        user_id SERIAL NOT NULL references users (id) on delete cascade,
        feed_id SERIAL NOT NULL references feeds (id) on delete cascade,
        created_at timestamptz (0) NULL DEFAULT now (),
        updated_at timestamptz (0) NULL
    );

-- +goose Down
DROP TABLE feed_user;