CREATE TABLE divisions (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name VARCHAR,
  phone VARCHAR,
  info TEXT,
  email VARCHAR,
  address VARCHAR,
  slug VARCHAR,
  floor_id uuid REFERENCES floors (id) ON UPDATE CASCADE ON DELETE CASCADE,
  entrance_id uuid REFERENCES entrances (id) ON UPDATE CASCADE ON DELETE CASCADE
);
