-- +goose Up
-- +goose StatementBegin
CREATE TABLE service_price
(
    id         INT PRIMARY KEY,
    created_at TIMESTAMP,
    price      DECIMAL(10, 2),
    FOREIGN KEY (id) REFERENCES service (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS service_price;
-- +goose StatementEnd
