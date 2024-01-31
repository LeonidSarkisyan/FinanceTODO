-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions (
      id SERIAL PRIMARY KEY,
      title VARCHAR(60),
      description VARCHAR(255),
      type VARCHAR(3) NOT NULL,
      value FLOAT NOT NULL,
      created_date_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_date_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      balance_id INT REFERENCES balances(id) ON DELETE CASCADE,
      sub_category_id INT REFERENCES sub_categories(id),
      user_id INT REFERENCES users(id)
);

CREATE OR REPLACE FUNCTION update_transaction_updated_date()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_date_time = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_transaction_updated_date
    BEFORE UPDATE ON transactions
    FOR EACH ROW
EXECUTE FUNCTION update_transaction_updated_date();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;
DROP TRIGGER IF EXISTS trigger_update_transaction_updated_date ON transactions;
DROP FUNCTION IF EXISTS update_transaction_updated_date;
DROP TABLE IF EXISTS transactions;
COMMIT;
-- +goose StatementEnd
