CREATE TABLE IF NOT EXISTS Profiles
(
    profiles_id        bigserial not null primary key,
    email              varchar   not null unique,
    nickname           varchar   not null unique,
    first_name         varchar   not null,
    second_name        varchar   not null,
    encrypted_password varchar   not null,
    avatar             varchar   not null,
    favorite_genre     text[] NOT NULL DEFAULT '{}'::text[]
);