-- +goose Up
CREATE TABLE items_categories(
    item_id UUID NOT NULL,
    category_id UUID NOT NULL,
    UNIQUE(item_id, category_id)
);

-- +goose Down
DROP TABLE items_categories;