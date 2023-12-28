CREATE TABLE parking_place (
    parking_place_id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255),
    address_id VARCHAR(50) REFERENCES address(address_id) ON DELETE CASCADE,
    account_id VARCHAR(50) REFERENCES account(account_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE parking_spot (
    parking_spot_id VARCHAR(50) PRIMARY KEY,
    parking_place_id VARCHAR(50) REFERENCES parking_place(parking_place_id) ON DELETE CASCADE,
    topic VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
