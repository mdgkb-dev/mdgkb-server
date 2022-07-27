CREATE TABLE paid_program_packages (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    price int,
    name varchar,
    paid_program_id uuid  REFERENCES paid_programs (id) ON UPDATE CASCADE ON DELETE CASCADE
);