-- Active: 1715341190893@@127.0.0.1@5432@Shop
CREATE ROLE not_auth_user WITH
    LOGIN
    NOSUPERUSER
    NOCREATEDB
    NOREPLICATION
    PASSWORD 'not_auth_user'
    CONNECTION LIMIT -1;

GRANT SELECT ON racket TO not_auth_user;
GRANT INSERT ON "user" TO not_auth_user;