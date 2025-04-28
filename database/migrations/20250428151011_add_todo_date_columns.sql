-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos
ADD COLUMN completed_at TIMESTAMP NULL,
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
ADD COLUMN deleted_at TIMESTAMP NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todos
DROP COLUMN completed_at,
DROP COLUMN created_at,
DROP COLUMN updated_at,
DROP COLUMN deleted_at;
-- +goose StatementEnd