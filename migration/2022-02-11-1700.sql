CREATE TABLE IF NOT EXISTS "user"
(
    id       bigserial    NOT NULL,
    name     varchar(255) NOT NULL,
    age      int,
    login    varchar(255) NOT NULL,
    email    varchar(255) NOT NULL,
    password varchar(255) NOT NULL
)