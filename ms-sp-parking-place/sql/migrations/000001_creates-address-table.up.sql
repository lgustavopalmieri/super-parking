-- init_schema.sql

CREATE TABLE address (
    address_id VARCHAR(50) PRIMARY KEY,
    parking_place_id VARCHAR(50),
    street VARCHAR(255),
    number INT,
    complement VARCHAR(255),
    district VARCHAR(255),
    uf VARCHAR(2),
    cep VARCHAR(10),
    lat VARCHAR(20),
    lng VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
