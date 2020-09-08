CREATE TABLE administrators (
    id serial primary key,
    username varchar(200),
    password varchar(200),
    salt varchar(200)
);


CREATE TABLE customers (
    id serial primary key,
    name varchar(200),
    phone varchar(200)
);


CREATE TABLE customer_addresses (
   id serial primary key,
   customer_id int,
   address varchar(200),
   zipcode varchar(5)
);
