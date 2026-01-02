-- +goose Up
-- +goose StatementBegin
-- Переименовываем таблицу title в sport_type
ALTER TABLE title RENAME TO sport_type;

-- Удаляем старую структуру (trainer_id, value) и создаем новую (только value)
-- Сначала удаляем данные и внешние ключи
ALTER TABLE sport_type DROP CONSTRAINT IF EXISTS title_trainer_id_fkey;
ALTER TABLE sport_type DROP COLUMN IF EXISTS trainer_id;

-- Переименовываем колонку value в name для ясности
ALTER TABLE sport_type RENAME COLUMN value TO name;

-- Делаем name NOT NULL и UNIQUE
ALTER TABLE sport_type ALTER COLUMN name SET NOT NULL;
ALTER TABLE sport_type ADD CONSTRAINT ux_sport_type_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Возвращаем обратно
ALTER TABLE sport_type DROP CONSTRAINT IF EXISTS ux_sport_type_name;
ALTER TABLE sport_type ALTER COLUMN name DROP NOT NULL;
ALTER TABLE sport_type RENAME COLUMN name TO value;
ALTER TABLE sport_type ADD COLUMN trainer_id INT;
ALTER TABLE sport_type ADD CONSTRAINT title_trainer_id_fkey FOREIGN KEY (trainer_id) REFERENCES trainer (id) ON DELETE CASCADE;
ALTER TABLE sport_type RENAME TO title;
-- +goose StatementEnd

