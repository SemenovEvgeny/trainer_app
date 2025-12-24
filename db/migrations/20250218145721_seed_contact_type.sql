-- +goose Up
-- +goose StatementBegin
INSERT INTO contact_type (value) VALUES
    ('phone'),
    ('email'),
    ('telegram'),
    ('whatsapp'),
    ('vk'),
    ('instagram'),
    ('max');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM contact_type WHERE value IN ('phone', 'email', 'telegram', 'whatsapp', 'vk', 'instagram', 'max');
-- +goose StatementEnd

