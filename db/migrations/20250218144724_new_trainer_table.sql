-- +goose Up
-- +goose StatementBegin
CREATE TABLE trainer
(
    id          SERIAL PRIMARY KEY,
    last_name   VARCHAR(255) NOT NULL,
    first_name  VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255),
    description TEXT,
    is_active   BOOL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS ux_trainer_names;
DROP TABLE IF EXISTS trainer;
-- +goose StatementEnd
