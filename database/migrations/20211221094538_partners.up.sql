CREATE TABLE partners
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    link varchar,
    image_id uuid REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    partner_type_id uuid REFERENCES partner_types (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
