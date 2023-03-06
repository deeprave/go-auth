CREATE TYPE cred_type AS ENUM (
    'plain',
    'hash',
    'totp'
);

CREATE TABLE IF NOT EXISTS credential (
    user_id bigint REFERENCES "user"(id),
    type cred_type DEFAULT 'plain',
    data jsonb NOT NULL DEFAULT '{}'::jsonb,
    UNIQUE (user_id, type)
);
