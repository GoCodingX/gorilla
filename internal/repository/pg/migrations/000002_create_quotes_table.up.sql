CREATE TABLE quotes
(
    id               UUID PRIMARY KEY             NOT NULL,
    author_id        UUID REFERENCES authors (id) NOT NULL,
    creator_username VARCHAR                      NOT NULL,
    text             VARCHAR                      NOT NULL,
    created_at       TIMESTAMPTZ                  NOT NULL,
    updated_at       TIMESTAMPTZ                  NOT NULL
);
