-- name: CreateItem :one
INSERT INTO items(id, created_at, updated_at, name, brand_id, category_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;