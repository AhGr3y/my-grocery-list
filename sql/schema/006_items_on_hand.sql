-- +goose Up
CREATE TABLE items_on_hand(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    quantity INTEGER NOT NULL,
    expiry_date TIMESTAMP,
    item_id UUID NOT NULL,
    CONSTRAINT fk_items
    FOREIGN KEY (item_id)
    REFERENCES items(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE items_on_hand;