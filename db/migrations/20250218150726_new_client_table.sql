-- +goose Up
-- +goose StatementBegin
CREATE TABLE client
(
    id          SERIAL PRIMARY KEY,
    last_name   VARCHAR(255) NOT NULL,
    first_name  VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255),
    description TEXT,
    is_active   BOOL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS client;
-- +goose StatementEnd
