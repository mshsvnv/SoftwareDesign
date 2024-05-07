-- Active: 1696673577093@@127.0.0.1@5432@Shop@public
create table if not exists "user" (
    id serial primary key,
    name text,
    surname text,
    email text,
    password text,
    role text
);

create table if not exists supplier (
    id serial primary key,
    name text,
    phone text,
    town text,
    email text
);

create table if not exists racket (
    id serial primary key,
    supplier_id int references supplier(id) on delete cascade,
    brand text,
    weight float,
    balance float,
    head_size float,
    price float,
    quantity int
);

create table if not exists cart (
    user_id int,
    quantity int,
    total_price float,
    primary key (user_id),
    foreign key (user_id) references "user"(id) on delete cascade
);

create table if not exists "order" (
    id serial primary key,
    user_id int references "user"(id) on delete cascade,
    delivery_date timestamp,
    address text,
    recepient_name text,
    status text,
    total_price float
);

create table if not exists order_racket (
    order_id int references "order"(id) on delete cascade,
    racket_id int references racket(id) on delete cascade,
    primary key (order_id, racket_id),
    quantity int
);

create table if not exists cart_racket (
    cart_id int references cart(user_id) on delete cascade,
    racket_id int references racket(id) on delete cascade,
    primary key (cart_id, racket_id),
    quantity int
);

create table if not exists feedback (
    racket_id int references racket(id) on delete cascade,
    user_id int references "user"(id) on delete cascade,
    primary key (user_id, racket_id),
    feedback text,
    rating float,
    date timestamp
);
