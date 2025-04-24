-- +goose Up
-- +goose StatementBegin
CREATE TABLE location
(
    id       SERIAL PRIMARY KEY,
    region   VARCHAR(255) NOT NULL,
    city     VARCHAR(255) NOT NULL,
    district VARCHAR(255) NOT NULL,
    street   VARCHAR(255) NOT NULL,
    house    VARCHAR(255) NOT NULL,
    text     VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS location;
-- +goose StatementEnd
