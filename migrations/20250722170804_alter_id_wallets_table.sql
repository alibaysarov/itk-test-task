-- +goose Up
-- +goose StatementBegin
ALTER TABLE wallets 
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE wallets 
ALTER COLUMN id DROP DEFAULT;
-- +goose StatementEnd
