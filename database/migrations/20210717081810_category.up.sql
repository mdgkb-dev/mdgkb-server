CREATE TABLE categories
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR
);
