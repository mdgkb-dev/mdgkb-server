CREATE TABLE fields
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    field_order INTEGER,
    form_id uuid  REFERENCES forms(id) ON UPDATE CASCADE ON DELETE CASCADE,
    value_type_id uuid references value_types
);
