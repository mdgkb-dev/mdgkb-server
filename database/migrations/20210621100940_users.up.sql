CREATE TABLE users
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    uuid uuid DEFAULT uuid_generate_v4() NOT NULL,
    email varchar unique,
    password varchar,
    phone varchar
);
