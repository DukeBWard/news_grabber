-- this comment below is the name of the sqlc go func that is generated with one or slice (many)
-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * from feed_follows WHERE user_id=$1;

-- exec just runs the command, returns no records thats why its just exec
-- $ is for params 
-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows WHERE id=$1 AND user_id=$2;