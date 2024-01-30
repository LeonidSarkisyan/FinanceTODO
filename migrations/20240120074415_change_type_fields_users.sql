-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS users
    ALTER COLUMN phone TYPE VARCHAR(255),
    ALTER COLUMN username TYPE VARCHAR(255),
    ALTER COLUMN email TYPE VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS users
    ALTER COLUMN phone TYPE TEXT,
    ALTER COLUMN username TYPE TEXT,
    ALTER COLUMN email TYPE TEXT;
-- +goose StatementEnd
