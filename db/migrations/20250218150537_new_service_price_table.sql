-- +goose Up
-- +goose StatementBegin
CREATE TABLE service_price
(
    id         SERIAL PRIMARY KEY,
    service_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    price      DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (service_id) REFERENCES service (id) ON DELETE CASCADE
);

-- Добавляем FK для price_id в service после создания service_price
ALTER TABLE service
    ADD CONSTRAINT fk_service_price 
    FOREIGN KEY (price_id) REFERENCES service_price (id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE service
    DROP CONSTRAINT IF EXISTS fk_service_price;
DROP TABLE IF EXISTS service_price;
-- +goose StatementEnd
