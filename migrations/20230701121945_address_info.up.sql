CREATE TABLE address_infos (
                          id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
                          region VARCHAR,
                          city VARCHAR,
                          street VARCHAR,
                          building VARCHAR,
                          flat VARCHAR,
                          zip int
);

alter table address_infos
    add column contact_info_id uuid REFERENCES contact_infos(id) ON UPDATE CASCADE ON DELETE CASCADE;