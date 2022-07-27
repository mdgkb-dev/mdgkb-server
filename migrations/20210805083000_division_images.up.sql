CREATE TABLE division_images (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  description varchar,
  division_id uuid REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE,
  file_info_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
