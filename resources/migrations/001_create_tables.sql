CREATE TABLE IF NOT EXISTS accounts (
  account_id BIGINT PRIMARY KEY,
  balance NUMERIC(20,10) NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
  transaction_id SERIAL PRIMARY KEY,
  source_account_id BIGINT NOT NULL REFERENCES accounts(account_id),
  destination_account_id BIGINT NOT NULL REFERENCES accounts(account_id),
  amount NUMERIC(20,10) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);
