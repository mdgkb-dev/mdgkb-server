CREATE TABLE departments (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    head_id uuid  REFERENCES heads (id) ON UPDATE CASCADE ON DELETE CASCADE,
    name varchar,
    is_division boolean,
    division_id uuid REFERENCES divisions (id)
);