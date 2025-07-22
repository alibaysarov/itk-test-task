-- +goose Up
CREATE TABLE IF NOT EXISTS wallets (
    id UUID PRIMARY KEY,
    amount INT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS wallets;
