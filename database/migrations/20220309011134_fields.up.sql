CREATE TABLE fields (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    required boolean default false,
    field_order INTEGER,
    form_id uuid REFERENCES forms(id) ON UPDATE CASCADE ON DELETE CASCADE,
    form_pattern_id uuid REFERENCES form_patterns(id) ON UPDATE CASCADE ON DELETE CASCADE,
    value_type_id uuid references value_types,
    form_value_id uuid  REFERENCES form_values (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    file_id uuid  REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);