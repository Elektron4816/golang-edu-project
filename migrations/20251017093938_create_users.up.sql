CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    email varchar(255) not null unique,
    password varchar(255) not null
);