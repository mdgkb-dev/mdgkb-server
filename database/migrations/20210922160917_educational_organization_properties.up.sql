CREATE TABLE educational_organization_properties (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    value varchar,
    educational_organization_property_type_id uuid REFERENCES educational_organization_property_types(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
