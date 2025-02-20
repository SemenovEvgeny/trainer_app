-- +goose Up
-- +goose StatementBegin
CREATE TABLE district
(
    id    SERIAL PRIMARY KEY,
    value VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS district;
-- +goose StatementEnd
