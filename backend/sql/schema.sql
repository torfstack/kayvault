CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE secrets
(
    id SERIAL PRIMARY KEY,
    value BYTEA NOT NULL,
    key TEXT NOT NULL,
    url TEXT NOT NULL,
    tags TEXT NOT NULL,
    user_id INTEGER NOT NULL
        REFERENCES users (id)
        ON DELETE CASCADE,
    secret_sharing INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
