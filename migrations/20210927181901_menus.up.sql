CREATE TABLE menus
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    top boolean,
    hide boolean,
    side boolean,
    menu_order int not null default 0,
    link VARCHAR,
    icon_id uuid ,
    page_id uuid
);
