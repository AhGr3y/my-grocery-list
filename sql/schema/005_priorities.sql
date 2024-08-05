-- +goose Up
CREATE TABLE priorities(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    item_id UUID NOT NULL,
    CONSTRAINT fk_items
    FOREIGN KEY (item_id)
    REFERENCES items(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE priorities;