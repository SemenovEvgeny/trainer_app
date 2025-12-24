-- +goose Up
-- +goose StatementBegin
CREATE TABLE service
(
    id          SERIAL PRIMARY KEY,
    trainer_id  INTEGER NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    location_id INTEGER,
    price_id    INTEGER,
    FOREIGN KEY (trainer_id) REFERENCES trainer (id) ON DELETE CASCADE,
    FOREIGN KEY (location_id) REFERENCES location (id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS service;
-- +goose StatementEnd
