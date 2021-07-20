CREATE TABLE divisions
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    number INTEGER,
    name VARCHAR,
    status VARCHAR,
    phone VARCHAR,
    email VARCHAR,
    address VARCHAR,
    slug VARCHAR,
    floor_id uuid  REFERENCES floors (id) ON UPDATE CASCADE ON DELETE CASCADE
);
