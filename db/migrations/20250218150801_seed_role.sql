-- +goose Up
-- +goose StatementBegin
INSERT INTO role (value) VALUES
    ('admin'),
    ('trainer'),
    ('sportsman');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM role WHERE value IN ('admin', 'trainer', 'sportsman');
-- +goose StatementEnd

