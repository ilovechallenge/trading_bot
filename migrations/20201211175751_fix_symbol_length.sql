-- +goose Up
ALTER TABLE trades MODIFY COLUMN symbol VARCHAR(9);
ALTER TABLE orders MODIFY COLUMN symbol VARCHAR(9);

-- +goose Down
ALTER TABLE trades MODIFY COLUMN symbol VARCHAR(8);
ALTER TABLE orders MODIFY COLUMN symbol VARCHAR(8);
