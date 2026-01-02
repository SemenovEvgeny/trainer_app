-- +goose Up
-- +goose StatementBegin
-- Связующая таблица для связи many-to-many между тренерами и видами спорта
CREATE TABLE trainer_sport
(
    id          SERIAL PRIMARY KEY,
    trainer_id  INTEGER NOT NULL,
    sport_id    INTEGER NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (trainer_id) REFERENCES trainer (id) ON DELETE CASCADE,
    FOREIGN KEY (sport_id) REFERENCES sport_type (id) ON DELETE CASCADE,
    CONSTRAINT ux_trainer_sport UNIQUE (trainer_id, sport_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS trainer_sport;
-- +goose StatementEnd

