CREATE TABLE carousels
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name    varchar,
    title    varchar
);
