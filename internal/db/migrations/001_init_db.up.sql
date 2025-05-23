CREATE TABLE IF NOT EXISTS todos (
     id SERIAL PRIMARY KEY,
     title TEXT    NOT NULL,
     completed BOOLEAN NOT NULL DEFAULT FALSE,
     created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);