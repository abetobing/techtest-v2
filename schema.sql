CREATE TABLE administrators (
    id int,
    username varchar(200),
    password varchar(200)
);

INSERT INTO administrators (id, username, password) VALUES
(1, 'admin', 'admin');

CREATE TABLE customers (
    id int,
    name varchar(200),
    phone varchar(200)
);


CREATE TABLE customer_addresses (
   id int,
   customer_id int,
   address varchar(200),
   zipcode varchar(5)
);
