CREATE TABLE authors
(
    id         UUID PRIMARY KEY NOT NULL,
    name       VARCHAR          NOT NULL,
    created_at TIMESTAMPTZ      NOT NULL NOT NULL,
    updated_at TIMESTAMPTZ      NOT NULL NOT NULL
);
