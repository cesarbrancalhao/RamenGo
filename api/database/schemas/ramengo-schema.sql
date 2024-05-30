USE ramengo;

CREATE TABLE IF NOT EXISTS broth (
    id INT AUTO_INCREMENT PRIMARY KEY,
    image_inactive TEXT,
    image_active TEXT,
    name VARCHAR(255),
    description TEXT,
    price INT
);

CREATE TABLE IF NOT EXISTS protein (
    id INT AUTO_INCREMENT PRIMARY KEY,
    image_inactive TEXT,
    image_active TEXT,
    name VARCHAR(255),
    description TEXT,
    price INT
);

CREATE TABLE IF NOT EXISTS order_response (
    id INT PRIMARY KEY,
    description TEXT,
    image TEXT
);