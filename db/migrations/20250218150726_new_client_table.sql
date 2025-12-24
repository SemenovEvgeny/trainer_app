-- +goose Up
-- +goose StatementBegin
CREATE TABLE sportsman
(
    id          SERIAL PRIMARY KEY,
    last_name   VARCHAR(255) NOT NULL,
    first_name  VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255),
    description TEXT,
    is_active   BOOL
);

-- Добавляем FK на sportsman в таблицу contact после создания sportsman
ALTER TABLE contact
    ADD CONSTRAINT contact_sportsman_id_fkey 
    FOREIGN KEY (sportsman_id) REFERENCES sportsman (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE contact DROP CONSTRAINT IF EXISTS contact_sportsman_id_fkey;
DROP TABLE IF EXISTS sportsman;
-- +goose StatementEnd
