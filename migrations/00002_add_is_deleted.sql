-- +goose Up
ALTER TABLE subscriptions ADD COLUMN is_deleted BOOLEAN DEFAULT false;

-- +goose Down
ALTER TABLE subscriptions DROP COLUMN is_deleted;
