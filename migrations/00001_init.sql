-- +goose Up
CREATE TABLE IF NOT EXISTS subscriptions(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(256) NOT NULL,
    price       INT NOT NULL,
    user_id     UUID NOT NULL,
    start_at    DATE NOT NULL,
    end_at      DATE NOT NULL,
    is_deleted  BOOLEAN DEFAULT false,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS subscriptions;
