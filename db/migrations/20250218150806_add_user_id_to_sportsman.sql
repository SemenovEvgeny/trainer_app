-- +goose Up
-- +goose StatementBegin
-- Добавляем user_id в таблицу sportsman для прямой связи один-к-одному
ALTER TABLE sportsman
    ADD COLUMN user_id INTEGER,
    ADD CONSTRAINT sportsman_user_id_fkey 
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL,
    ADD CONSTRAINT ux_sportsman_user_id UNIQUE (user_id);

-- Удаляем связующую таблицу user_sportsman, так как теперь связь прямая
DROP TABLE IF EXISTS user_sportsman;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Восстанавливаем связующую таблицу
CREATE TABLE IF NOT EXISTS user_sportsman
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER NOT NULL,
    sportsman_id INTEGER NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (sportsman_id) REFERENCES sportsman (id) ON DELETE CASCADE,
    CONSTRAINT ux_user_sportsman UNIQUE (user_id, sportsman_id)
);

-- Удаляем user_id из sportsman
ALTER TABLE sportsman
    DROP CONSTRAINT IF EXISTS ux_sportsman_user_id,
    DROP CONSTRAINT IF EXISTS sportsman_user_id_fkey,
    DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd

