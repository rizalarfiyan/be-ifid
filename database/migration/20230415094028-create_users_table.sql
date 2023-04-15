
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	email VARCHAR (255) UNIQUE NOT NULL,
    first_name VARCHAR (100),
    last_name VARCHAR (255),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
INSERT INTO users(email, first_name, last_name) VALUES ('rizal.arfiyan.23@gmail.com', 'Rizal', 'Arfiyan');

-- +migrate Down
DROP TABLE users;
