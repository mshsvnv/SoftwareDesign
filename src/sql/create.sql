-- Active: 1696673577093@@127.0.0.1@5432@Shop
create table if not exists "user" (
    id text primary key,
    email text,
    password text,
    role text
);

create table if not exists racket (
    id text primary key,
    brand text,
    price float,
    quantity int
);

create table if not exists cart (
    id text primary key,
    user_id text references "user"(id) on delete cascade,
    rackets_quantity int
);

create table if not exists "order" (
    id text primary key,
    user_id text references "user"(id) on delete cascade,
    creation_date timestamp,
    delivery_date timestamp,
    address text,
    status text
);

create table if not exists payment (
    order_id text references "order"(id) on delete cascade,
    status text,
    pay_date timestamp
);

create table if not exists order_racket (
    order_id text references "order"(id) on delete cascade,
    racket_id text references racket(id) on delete cascade,
    quantity int
);

create table if not exists cart_racket (
    order_id text references cart(id) on delete cascade,
    cart_id text references racket(id) on delete cascade,
    quantity int
);