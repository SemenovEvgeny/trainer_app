-- +goose Up
-- +goose StatementBegin
CREATE TABLE region
(
    id    SERIAL PRIMARY KEY,
    value VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS region;
-- +goose StatementEnd
