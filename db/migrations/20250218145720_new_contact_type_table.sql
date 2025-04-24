-- +goose Up
-- +goose StatementBegin
CREATE TABLE contact_type
(
    id    SERIAL PRIMARY KEY,
    value VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS contact_type;
-- +goose StatementEnd
