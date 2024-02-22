CREATE TABLE IF NOT EXISTS accounts (
  id INTEGER PRIMARY KEY,
  account_limit INTEGER NOT NULL,
  balance INTEGER NOT NULL DEFAULT 0
);
CREATE INDEX accounts_id_idx ON "accounts" USING HASH(id);

CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  account_id INTEGER NOT NULL,
  amount INTEGER NOT NULL,
  operation CHAR(1) NOT NULL,
  description VARCHAR(10) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE INDEX transactions_id_idx ON "transactions" USING HASH(id);
CREATE INDEX transactions_account_id_idx ON "transactions" USING HASH(account_id);

CREATE OR REPLACE PROCEDURE create_transaction(
  account_id INTEGER,
  amount INTEGER,
  operation CHAR(1),
  description VARCHAR(10),
  inout balance_updated integer default null,
  inout limit_updated integer default null
)
LANGUAGE plpgsql
AS $$

BEGIN
  UPDATE accounts SET balance = balance + amount WHERE id = account_id and balance + amount >= - account_limit
  RETURNING balance, account_limit into balance_updated, limit_updated;
  IF balance_updated IS NULL OR limit_updated IS NULL THEN RETURN; END IF;

  COMMIT;

  INSERT INTO transactions (account_id, amount, operation, description)
  VALUES (account_id, ABS(amount), operation, description);
END;
$$;

DO $$
BEGIN
  INSERT INTO accounts (id, account_limit)
  VALUES
    (1, 100000),
    (2, 80000),
    (3, 1000000),
    (4, 10000000),
    (5, 500000);
END;
$$