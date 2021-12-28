CREATE TABLE paid_program_services_groups (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    home boolean,
    group_order int,
    paid_program_package_id uuid  REFERENCES paid_program_packages (id) ON UPDATE CASCADE ON DELETE CASCADE
);