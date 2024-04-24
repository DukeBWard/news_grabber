-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
-- This is SQLC stuff.  Each of these are params for the function
VALUES ($1, $2, $3, $4)
RETURNING *;