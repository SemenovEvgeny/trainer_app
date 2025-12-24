-- +goose Up
-- +goose StatementBegin
CREATE TABLE achievement
(
    id         SERIAL PRIMARY KEY,
    trainer_id INT  NOT NULL,
    value      TEXT NOT NULL,
    FOREIGN KEY (trainer_id) REFERENCES trainer (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS achievement;
-- +goose StatementEnd
