-- +goose Up
-- +goose StatementBegin
ALTER TABLE title
    RENAME COLUMN titles TO value;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE title
    RENAME COLUMN value TO titles;
-- +goose StatementEnd
