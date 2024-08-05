// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: items.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items(id, created_at, updated_at, name, brand_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, name, brand_id
`

type CreateItemParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	BrandID   uuid.UUID
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.BrandID,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.BrandID,
	)
	return i, err
}