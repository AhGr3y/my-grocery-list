-- name: CreatePriority :one
INSERT INTO priorities (id, created_at, updated_at, name, item_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;