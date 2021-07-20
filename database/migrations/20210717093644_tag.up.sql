CREATE TABLE tags
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    color VARCHAR,
    label VARCHAR
);
