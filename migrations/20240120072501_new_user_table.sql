-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    phone TEXT NOT NULL UNIQUE,
    username TEXT,
    email TEXT,
    password TEXT NOT NULL,
    created_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_datetime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_datetime = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_users
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_datetime();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_users ON users;
DROP FUNCTION IF EXISTS update_updated_datetime();
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
