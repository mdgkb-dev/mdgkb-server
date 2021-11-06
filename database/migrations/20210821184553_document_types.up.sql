create table document_types
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar not null,
    single_scan boolean,
    scan_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
