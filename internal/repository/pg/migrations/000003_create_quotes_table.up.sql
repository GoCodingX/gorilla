CREATE TABLE quotes
(
    id              UUID PRIMARY KEY             NOT NULL,
    author_id       UUID REFERENCES authors (id) NOT NULL,
    creator_user_id UUID REFERENCES users (id)   NOT NULL,
    text            VARCHAR                      NOT NULL,
    created_at      TIMESTAMPTZ                  NOT NULL,
    updated_at      TIMESTAMPTZ                  NOT NULL
);
