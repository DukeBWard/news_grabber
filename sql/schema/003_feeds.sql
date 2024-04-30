-- +goose Up
-- Up Migration

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    -- will get an error if trying to make a uuid that is NOT in users table
    -- the on delete cascade is self explainatory, if user is deleted so are their feeds
    user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
-- Down Migration: Undo the up

DROP TABLE feeds;