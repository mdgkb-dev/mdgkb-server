CREATE TABLE centers (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name VARCHAR,
  info TEXT,
  address VARCHAR,
  slug VARCHAR
);
