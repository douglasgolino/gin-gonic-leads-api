CREATE TABLE leads (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
