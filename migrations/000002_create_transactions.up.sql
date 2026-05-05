CREATE TABLE transactions (
    id          SERIAL PRIMARY KEY,
    user_id     INT NOT NULL REFERENCES users(id),
    description TEXT NOT NULL,
    amount      NUMERIC(10,2) NOT NULL,
    type        TEXT NOT NULL CHECK (type IN ('income', 'expense')),
    category    TEXT,
    created_at  TIMESTAMP DEFAULT NOW()
);