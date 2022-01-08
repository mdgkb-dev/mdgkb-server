CREATE TABLE doctors
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    human_id uuid  REFERENCES humans (id) ON UPDATE CASCADE ON DELETE CASCADE,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    show boolean default false,
    education         VARCHAR,
    schedule          VARCHAR,
    position          VARCHAR,
    tags              VARCHAR,
    file_info_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
