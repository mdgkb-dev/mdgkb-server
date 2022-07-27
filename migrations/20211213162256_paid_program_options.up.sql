CREATE TABLE paid_program_options (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    option_default boolean,
    item_order integer,
    paid_program_options_group_id uuid  REFERENCES paid_program_options_groups (id) ON UPDATE CASCADE ON DELETE CASCADE
);
