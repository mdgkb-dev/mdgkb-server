CREATE TABLE pages
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    title varchar,
    content text,
    slug VARCHAR,
    link VARCHAR
);
