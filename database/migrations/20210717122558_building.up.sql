CREATE TABLE buildings
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    number VARCHAR,
    address          VARCHAR
);
