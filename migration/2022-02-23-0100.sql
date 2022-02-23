CREATE TABLE IF NOT EXISTS token
(
    user_id    bigserial references "user" (id),
    token      varchar(511) NOT NULL,
    is_valid   bool,
    created_at timestamptz  NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS authorization_attempts
(
    user_id    bigserial references "user" (id),
    is_success bool,
    sign_in    timestamptz  NOT NULL DEFAULT (now()),
    ip         varchar(255) NOT NULL
);