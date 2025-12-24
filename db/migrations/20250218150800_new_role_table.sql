-- +goose Up
-- +goose StatementBegin
CREATE TABLE role
(
    id    SERIAL PRIMARY KEY,
    value VARCHAR(50) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role;
-- +goose StatementEnd

