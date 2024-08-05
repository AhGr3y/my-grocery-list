-- name: CreateItemOnHand :one
INSERT INTO items_on_hand(id, created_at, updated_at, quantity, expiry_date, item_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;