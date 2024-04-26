-- +goose Up
-- Up Migration

ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    -- generating random bytes and casting to a byte array
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
-- Down Migration: Undo the up
ALTER TABLE users DROP COLUMN api_key;