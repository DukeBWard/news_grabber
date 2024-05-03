-- +goose Up
-- Up Migration

CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    -- this combination needs to be unique
    UNIQUE(user_id, feed_id)
);

-- +goose Down
-- Down Migration: Undo the up

DROP TABLE feed_follows;