-- +goose Up
-- +goose StatementBegin
CREATE TABLE service
(
    id          SERIAL PRIMARY KEY,
    trainer_id  INT,
    name        VARCHAR(255),
    price_id    INT,
    description TEXT,
    location_id INT,
    FOREIGN KEY (trainer_id) REFERENCES trainer (id) ON DELETE CASCADE,
    FOREIGN KEY (price_id) REFERENCES service_price (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS service;
-- +goose StatementEnd
