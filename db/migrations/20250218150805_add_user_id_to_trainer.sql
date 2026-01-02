-- +goose Up
-- +goose StatementBegin
-- Добавляем user_id в таблицу trainer для прямой связи один-к-одному
ALTER TABLE trainer
    ADD COLUMN user_id INTEGER,
    ADD CONSTRAINT trainer_user_id_fkey 
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL,
    ADD CONSTRAINT ux_trainer_user_id UNIQUE (user_id);

-- Удаляем связующую таблицу user_trainer, так как теперь связь прямая
DROP TABLE IF EXISTS user_trainer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Восстанавливаем связующую таблицу
CREATE TABLE IF NOT EXISTS user_trainer
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER NOT NULL,
    trainer_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (trainer_id) REFERENCES trainer (id) ON DELETE CASCADE,
    CONSTRAINT ux_user_trainer UNIQUE (user_id, trainer_id)
);

-- Удаляем user_id из trainer
ALTER TABLE trainer
    DROP CONSTRAINT IF EXISTS ux_trainer_user_id,
    DROP CONSTRAINT IF EXISTS trainer_user_id_fkey,
    DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd

