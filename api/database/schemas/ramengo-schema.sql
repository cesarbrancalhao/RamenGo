CREATE TABLE IF NOT EXISTS broth (
    id VARCHAR(255) PRIMARY KEY,
    image_inactive TEXT,
    image_active TEXT,
    name VARCHAR(255),
    description TEXT,
    price INT
);

CREATE TABLE IF NOT EXISTS protein (
    id VARCHAR(255) PRIMARY KEY,
    image_inactive TEXT,
    image_active TEXT,
    name VARCHAR(255),
    description TEXT,
    price INT
);

CREATE TABLE IF NOT EXISTS order_response (
    id VARCHAR(255) PRIMARY KEY,
    description TEXT,
    image TEXT
);