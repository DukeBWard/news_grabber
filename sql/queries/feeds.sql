-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
-- This is SQLC stuff.  Each of these are params for the function
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * from feeds;

-- name: GetNextFeedToFetch :one
SELECT * from feeds 
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT 1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
set last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
-- returning * basically returns all of the tables that were affect like a select *
-- for example, if its a delete it will return the rows that are deleted
RETURNING *;