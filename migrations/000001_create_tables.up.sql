BEGIN;
CREATE TABLE users (
id SERIAL PRIMARY KEY,
username VARCHAR(64) UNIQUE NOT NULL,
email VARCHAR(255) UNIQUE NOT NULL,
password VARCHAR(64) NOT NULL,
is_admin BOOLEAN NOT NULL DEFAULT FALSE,
is_banned BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE apps (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL
);

CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL
);

CREATE TABLE refresh_tokens (
    token VARCHAR(128) PRIMARY KEY,
    app_id INTEGER NOT NULL,
    client_id INTEGER NOT NULL,
    expires_at BIGINT NOT NULL,
    UNIQUE (app_id, client_id),
    CONSTRAINT fk_app
        FOREIGN KEY (app_id) 
        REFERENCES apps(id) 
        ON DELETE CASCADE,

    CONSTRAINT fk_client
        FOREIGN KEY (client_id)
        REFERENCES clients(id)
        ON DELETE CASCADE
    
);

INSERT INTO apps(name) VALUES ('clown-space');
INSERT INTO clients(name) VALUES ('web'), ('mobile');

END;