CREATE TABLE IF NOT EXISTS "users"
(
    "id"             serial PRIMARY KEY,
    "name"           VARCHAR(50)         NOT NULL,
    "email"          VARCHAR(255)        NOT NULL,
    "login"          VARCHAR(255)        NOT NULL,
    "password"       VARCHAR(255)        NOT NULL
);

