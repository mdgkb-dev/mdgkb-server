CREATE TABLE hospitalizations (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    paid boolean
);