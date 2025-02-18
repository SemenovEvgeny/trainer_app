-- +goose Up
-- +goose StatementBegin
CREATE TABLE title
(
    id         SERIAL PRIMARY KEY,
    trainer_id INT,
    titles     TEXT,
    FOREIGN KEY (trainer_id) REFERENCES trainer (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS trainer;
-- +goose StatementEnd
