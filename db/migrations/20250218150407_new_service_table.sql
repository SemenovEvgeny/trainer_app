-- +goose Up
-- +goose StatementBegin
CREATE TABLE service
(
    id          SERIAL PRIMARY KEY,
    trainer_id  INT,
    name        VARCHAR(255),
    description TEXT,
    price       DECIMAL(10, 2),
    locations   TEXT,
    FOREIGN KEY (trainer_id) REFERENCES trainer (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS service;
-- +goose StatementEnd
