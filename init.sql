CREATE TABLE IF NOT EXISTS wallets (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL UNIQUE,
    balance DECIMAL(10,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_wallets_user_id ON wallets(user_id);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    wallet_id INTEGER NOT NULL,
    transaction_type VARCHAR(20) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    balance_before DECIMAL(10,2),
    balance_after DECIMAL(10,2),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);

CREATE INDEX IF NOT EXISTS idx_transactions_wallet_id ON transactions(wallet_id);

-- dummy wallets
INSERT INTO wallets (user_id, balance, created_at, updated_at) VALUES
('user123', 5000.00, NOW(), NOW()),
('user456', 3000.00, NOW(), NOW()),
('user789', 10000.00, NOW(), NOW()),
('alice', 2500.50, NOW(), NOW()),
('bob', 7500.75, NOW(), NOW());

-- dunmmy data
INSERT INTO transactions (wallet_id, transaction_type, amount, balance_before, balance_after, description, created_at) VALUES
(1, 'deposit', 5000.00, 0.00, 5000.00, 'Initial deposit', NOW()),
(2, 'deposit', 3000.00, 0.00, 3000.00, 'Initial deposit', NOW()),
(3, 'deposit', 10000.00, 0.00, 10000.00, 'Initial deposit', NOW()),
(4, 'deposit', 2500.50, 0.00, 2500.50, 'Initial deposit', NOW()),
(5, 'deposit', 7500.75, 0.00, 7500.75, 'Initial deposit', NOW());

--DROP TABLE IF EXISTS transactions CASCADE;
--DROP TABLE IF EXISTS wallets CASCADE;