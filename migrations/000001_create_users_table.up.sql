BEGIN;
CREATE TABLE users (
id serial PRIMARY KEY,
username VARCHAR(64) UNIQUE NOT NULL,
email VARCHAR(255) UNIQUE NOT NULL,
password VARCHAR(64) NOT NULL
);
END;