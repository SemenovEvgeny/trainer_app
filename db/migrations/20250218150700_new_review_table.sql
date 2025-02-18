-- +goose Up
-- +goose StatementBegin
CREATE TABLE review
(
    id         SERIAL PRIMARY KEY,
    service_id INT,
    score      INT CHECK (score BETWEEN 1 AND 5),
    comment    TEXT,
    FOREIGN KEY (service_id) REFERENCES service (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS review;
-- +goose StatementEnd
