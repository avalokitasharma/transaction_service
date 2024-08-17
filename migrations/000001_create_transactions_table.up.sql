CREATE TABLE IF NOT EXISTS transactions (
    id BIGINT PRIMARY KEY,
    amount DOUBLE PRECISION NOT NULL,
    type VARCHAR(255) NOT NULL,
    parent_id BIGINT REFERENCES transactions(id)
);