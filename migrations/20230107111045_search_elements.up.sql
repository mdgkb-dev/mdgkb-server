create extension pg_trgm;

CREATE TABLE search_elements
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    key varchar ,
    label varchar,
    description varchar,
    search_column tsvector,
    value varchar
);