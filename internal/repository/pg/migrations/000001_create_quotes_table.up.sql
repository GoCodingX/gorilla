CREATE TABLE quotes
(
    id          UUID PRIMARY KEY,
    description VARCHAR,
    created_at  TIMESTAMPTZ NOT NULL,
    updated_at  TIMESTAMPTZ NOT NULL
);
