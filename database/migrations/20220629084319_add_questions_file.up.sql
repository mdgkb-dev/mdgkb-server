alter table questions
add column file_id uuid REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
