-- +goose Up
-- +goose StatementBegin
CREATE TABLE location
(
    id          SERIAL PRIMARY KEY,
    region_id   INT NOT NULL,
    city_id     INT NOT NULL,
    district_id INT NOT NULL,
    street_id   INT NOT NULL,
    text        VARCHAR(255),
    FOREIGN KEY (region_id) REFERENCES region (id),
    FOREIGN KEY (city_id) REFERENCES city (id),
    FOREIGN KEY (district_id) REFERENCES district (id),
    FOREIGN KEY (street_id) REFERENCES street (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS location;
-- +goose StatementEnd
