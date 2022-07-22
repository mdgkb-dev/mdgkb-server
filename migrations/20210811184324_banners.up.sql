CREATE TABLE banners (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name varchar,
  link varchar,
  list_number integer,
  file_info_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
