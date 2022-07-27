CREATE TABLE paid_programs (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    description varchar,
    program_order int,
    paid_programs_group_id uuid  REFERENCES paid_programs_groups (id) ON UPDATE CASCADE ON DELETE CASCADE
);