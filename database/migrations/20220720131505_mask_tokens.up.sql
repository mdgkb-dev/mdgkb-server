CREATE TABLE mask_tokes (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    pattern VARCHAR,
    uppercase boolean default false,

    field_id uuid  REFERENCES fields (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
