-- init.sql

CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    request TEXT NOT NULL,
    responsecode INT NOT NULL,
    response TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
