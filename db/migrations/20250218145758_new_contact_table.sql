-- +goose Up
-- +goose StatementBegin
CREATE TABLE contact
(
    id          SERIAL PRIMARY KEY,
    trainer_id  INTEGER,
    sportsman_id INTEGER,
    type_id     INTEGER          NOT NULL,
    contact     VARCHAR(255) NOT NULL,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (trainer_id) REFERENCES trainer (id) ON DELETE CASCADE,
    FOREIGN KEY (type_id) REFERENCES contact_type (id),
    CONSTRAINT chk_contact_owner CHECK (
        (trainer_id IS NOT NULL AND sportsman_id IS NULL) OR 
        (trainer_id IS NULL AND sportsman_id IS NOT NULL)
    )
);

-- FK на sportsman будет добавлен после создания таблицы sportsman

-- Уникальный индекс для контактов (работает для тренеров и спортсменов)
-- Используем COALESCE для объединения trainer_id и sportsman_id
CREATE UNIQUE INDEX ux_contact_owner_type_contact 
    ON contact (COALESCE(trainer_id, sportsman_id), type_id, contact);

-- Создаем функцию для автоматического обновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ language 'plpgsql';

-- Создаем триггер для автоматического обновления updated_at при изменении записи
CREATE TRIGGER update_contact_updated_at
    BEFORE UPDATE ON contact
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_contact_updated_at ON contact;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP INDEX IF EXISTS ux_contact_owner_type_contact;
DROP TABLE IF EXISTS contact;
-- +goose StatementEnd
