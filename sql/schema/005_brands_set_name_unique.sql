-- +goose Up
ALTER TABLE brands
ADD CONSTRAINT unique_name UNIQUE (name);

-- +goose Down
ALTER TABLE brands
DROP CONSTRAINT unique_name;