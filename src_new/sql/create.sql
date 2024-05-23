create table if not exists "user" (
    id serial primary key,
    name text,
    surname text,
    email text unique,
    password text,
    subscription boolean,
    role text
);

create table if not exists supplier (
    id serial primary key,
    email text unique,
    name text,
    phone text,
    town text
);

create table if not exists racket (
    id serial primary key,
    supplier_id int,
    brand text,
    weight float,
    balance float,
    head_size float,
    avaliable boolean,
    price float,
    quantity int,
    foreign key (supplier_id) references supplier(id) on delete cascade
);

create table if not exists cart (
    user_id int unique,
    quantity int,
    total_price float,
    primary key (user_id),
    foreign key (user_id) references "user"(id) on delete cascade
);

create table if not exists "order" (
    id serial primary key,
    user_id int,
    status text,
    total_price float,
    creation_date timestamp,
    foreign key (user_id) references "user"(id) on delete cascade
);

create table if not exists delivery (
    order_id int,
    delivery_date timestamp,
    address text,
    recepient_name text,
    foreign key (order_id) references "order"(id) on delete cascade
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
