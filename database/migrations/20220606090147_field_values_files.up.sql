CREATE TABLE field_values_files (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    field_value_id uuid  REFERENCES field_values (id) ON UPDATE CASCADE ON DELETE CASCADE,
    file_info_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
