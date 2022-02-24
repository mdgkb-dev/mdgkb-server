CREATE TABLE educational_managers (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    doctor_id uuid REFERENCES doctors (id),
    role varchar,
    educational_manager_order int
);
