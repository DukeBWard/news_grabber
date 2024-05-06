-- this comment below is the name of the sqlc go func that is generated with one or slice (many)
-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
-- This is SQLC stuff.  Each of these are params for the function
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
