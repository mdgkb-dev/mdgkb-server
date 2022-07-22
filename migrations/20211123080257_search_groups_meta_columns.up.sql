CREATE TABLE search_group_meta_columns
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    label varchar,
    name varchar,
    search_group_id uuid REFERENCES search_groups(id) ON UPDATE CASCADE ON DELETE CASCADE
);
