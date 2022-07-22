CREATE TABLE public_document_types (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    route_anchor VARCHAR,
    description varchar,
    name VARCHAR
);