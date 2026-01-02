-- +goose Up
-- +goose StatementBegin
CREATE TABLE client_service
(
    id         SERIAL PRIMARY KEY,
    client_id  INTEGER NOT NULL,
    service_id INTEGER NOT NULL,
    price_id   INTEGER NOT NULL,
    date       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active  BOOL DEFAULT TRUE,
    FOREIGN KEY (client_id) REFERENCES sportsman (id) ON DELETE CASCADE,
    FOREIGN KEY (service_id) REFERENCES service (id) ON DELETE CASCADE,
    FOREIGN KEY (price_id) REFERENCES service_price (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS client_service;
-- +goose StatementEnd
