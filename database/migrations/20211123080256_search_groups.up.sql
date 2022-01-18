CREATE TABLE search_groups
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    key varchar,
    search_group_order integer,
    route varchar,
    search_group_table varchar,
    search_column varchar,
    label varchar,
    label_column varchar,
    value_column varchar
);
