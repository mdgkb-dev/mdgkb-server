CREATE TABLE users
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    email varchar,
    password varchar
);
