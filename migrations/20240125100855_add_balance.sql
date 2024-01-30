-- +goose Up
-- +goose StatementBegin
CREATE TABLE balances (
    id SERIAL PRIMARY KEY,
    title VARCHAR(16),
    type VARCHAR,
    value DOUBLE PRECISION,
    currency VARCHAR,
    created_date_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT REFERENCES users(id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION update_balances_updated_time()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_date_time = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_balances_trigger
    BEFORE UPDATE ON balances
    FOR EACH ROW
EXECUTE FUNCTION update_balances_updated_time();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_balances_trigger ON balances;
DROP FUNCTION IF EXISTS update_balances_updated_time();
DROP TABLE IF EXISTS balances;
-- +goose StatementEnd
