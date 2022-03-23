CREATE TABLE users
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    uuid uuid DEFAULT NULL,
    email varchar unique,
    password varchar,
    phone varchar
);
