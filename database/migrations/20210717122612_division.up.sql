CREATE TABLE divisions (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name VARCHAR,
  show boolean default false,
  info TEXT,
  address VARCHAR,
  slug VARCHAR,
  show_common_visiting_rules boolean default true,
  floor_id uuid REFERENCES floors (id) ON UPDATE CASCADE ON DELETE CASCADE,
  entrance_id uuid REFERENCES entrances (id) ON UPDATE CASCADE ON DELETE CASCADE,
  contact_info_id uuid references contact_infos(id) on update cascade on delete cascade
);
