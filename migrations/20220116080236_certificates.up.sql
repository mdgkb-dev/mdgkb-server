CREATE TABLE certificates
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    doctor_id uuid  REFERENCES doctors (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    scan_id uuid  REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    description varchar
);