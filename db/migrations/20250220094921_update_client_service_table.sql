-- +goose Up
-- +goose StatementBegin
ALTER TABLE client_service
    ADD COLUMN price_id INTEGER REFERENCES service_price (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE client_service
DROP
COLUMN price_id;
-- +goose StatementEnd
