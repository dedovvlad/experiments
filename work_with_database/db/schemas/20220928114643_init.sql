-- +goose Up
CREATE TABLE passports (
    id     character varying(128) PRIMARY KEY,
    series character varying(128) NOT NULL,
    number character varying(128) NOT NULL
);

-- +goose Down
DROP TABLE passports
