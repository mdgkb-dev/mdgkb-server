CREATE TABLE paid_programs_groups (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    group_order int
);