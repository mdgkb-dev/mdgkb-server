CREATE TABLE positions
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    show boolean default true,
    item_order int default 0
);
