-- +goose Up
CREATE TABLE items(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    brand_id UUID NOT NULL,
    CONSTRAINT fk_brands
    FOREIGN KEY (brand_id)
    REFERENCES brands(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE items;