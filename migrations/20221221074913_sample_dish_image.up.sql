alter table dishes_samples
add column image_id uuid REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
