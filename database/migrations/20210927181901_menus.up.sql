CREATE TABLE menus
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    top boolean,
    side boolean,
    link VARCHAR,
    page_id uuid
);
