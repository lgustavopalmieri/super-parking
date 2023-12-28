CREATE TABLE account (
    account_id VARCHAR(50) PRIMARY KEY,
    parking_place_id VARCHAR(50),
    balance FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
