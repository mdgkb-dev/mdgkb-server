CREATE TABLE form_value_files
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    form_value_id  uuid  REFERENCES form_values(id) ON UPDATE CASCADE ON DELETE CASCADE,
    file_id uuid  REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
