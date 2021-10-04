CREATE TABLE sub_sub_menus
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    link VARCHAR,
    page_id uuid ,
    icon_id uuid ,
    sub_menu_id uuid  REFERENCES sub_menus (id) ON UPDATE CASCADE ON DELETE CASCADE
);
