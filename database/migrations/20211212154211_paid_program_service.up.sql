CREATE TABLE paid_program_services (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    quantity varchar,
    service_order int,
    paid_program_services_group_id uuid  REFERENCES paid_program_services_groups (id) ON UPDATE CASCADE ON DELETE CASCADE
);