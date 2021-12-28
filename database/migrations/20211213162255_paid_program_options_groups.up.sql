CREATE TABLE paid_program_options_groups (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    item_order integer,
    paid_program_id uuid  REFERENCES paid_programs (id) ON UPDATE CASCADE ON DELETE CASCADE
);

