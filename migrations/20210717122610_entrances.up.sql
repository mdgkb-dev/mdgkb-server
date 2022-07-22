CREATE TABLE entrances (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  number INTEGER,
  building_id uuid REFERENCES buildings (id) ON UPDATE CASCADE ON DELETE CASCADE
);
