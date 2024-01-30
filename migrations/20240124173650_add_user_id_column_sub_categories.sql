-- +goose Up
-- +goose StatementBegin
ALTER TABLE sub_categories
    ADD COLUMN user_id INT REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sub_categories
    DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd
