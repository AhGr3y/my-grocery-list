-- name: CreateItemCategoryLink :one
INSERT INTO items_categories(item_id, category_id)
VALUES ($1, $2)
RETURNING *;