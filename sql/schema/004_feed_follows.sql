-- +goose Up
-- Up Migration

CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id NOT NULL
);

-- +goose Down
-- Down Migration: Undo the up

DROP TABLE feed_follows;