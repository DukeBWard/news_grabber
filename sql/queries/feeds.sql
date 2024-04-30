-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
-- This is SQLC stuff.  Each of these are params for the function
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
