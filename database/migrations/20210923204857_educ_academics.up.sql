CREATE TABLE educational_organization_academics (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    doctor_id uuid REFERENCES doctors (id)
);