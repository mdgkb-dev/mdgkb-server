CREATE TABLE sub_menus
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    link VARCHAR,
    page_id uuid  ,
    menu_id uuid  REFERENCES menus (id) ON UPDATE CASCADE ON DELETE CASCADE
);
