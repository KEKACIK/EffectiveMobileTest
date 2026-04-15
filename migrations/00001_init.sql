-- +goose Up
CREATE TABLE IF NOT EXISTS subscribers(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(256) NOT NULL,
    price       INT,
    user_id     UUID NOT NULL,
    start_at    DATE NOT NULL,
    end_at      DATE NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS subscribers;
