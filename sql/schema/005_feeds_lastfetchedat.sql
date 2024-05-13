-- +goose Up
-- Up Migration

ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
-- Down Migration: Undo the up
ALTER TABLE feeds DROP COLUMN last_fetched_at;