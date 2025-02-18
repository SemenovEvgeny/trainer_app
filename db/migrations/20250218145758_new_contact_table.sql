-- +goose Up
-- +goose StatementBegin
CREATE TABLE contact
(
    id         SERIAL PRIMARY KEY,
    trainer_id INT          NOT NULL,
    type_id    INT          NOT NULL,
    contact    VARCHAR(255) NOT NULL,
    FOREIGN KEY (trainer_id) REFERENCES trainer (id),
    FOREIGN KEY (type_id) REFERENCES contact_type (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS trainer;
-- +goose StatementEnd
