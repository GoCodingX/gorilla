CREATE TABLE users
(
    id         UUID PRIMARY KEY NOT NULL,
    username   VARCHAR          NOT NULL,
    created_at TIMESTAMPTZ      NOT NULL,
    updated_at TIMESTAMPTZ      NOT NULL
);
