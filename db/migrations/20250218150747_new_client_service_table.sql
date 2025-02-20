-- +goose Up
-- +goose StatementBegin
CREATE TABLE client_service
(
    id         SERIAL PRIMARY KEY,
    client_id  INT,
    service_id INT,
    price_id   INT,
    date       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active  BOOL,
    FOREIGN KEY (client_id) REFERENCES client (id) ON DELETE CASCADE,
    FOREIGN KEY (service_id) REFERENCES service (id) ON DELETE CASCADE,
    FOREIGN KEY (price_id) REFERENCES service_price (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS client_service;
-- +goose StatementEnd
