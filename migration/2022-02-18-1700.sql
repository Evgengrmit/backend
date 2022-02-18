CREATE TABLE IF NOT EXISTS "account"
(
    id       bigserial NOT NULL,
    amount   bigint    NOT NULL DEFAULT 0,
    currency int       NOT NULL,
    user_id  bigserial references "user" (id)
)