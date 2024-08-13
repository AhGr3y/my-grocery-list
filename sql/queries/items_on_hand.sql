-- name: StoreItem :one
INSERT INTO items_on_hand(id, created_at, updated_at, item_id, quantity, expiry_date, priority)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetItemOnHand :one
SELECT * FROM items_on_hand
WHERE item_id = $1;