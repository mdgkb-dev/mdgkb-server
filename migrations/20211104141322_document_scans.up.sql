create table documents_scans
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_2e1aa55eac1947ddf3221506edb"
            primary key,
    document_id uuid not null references documents ON UPDATE CASCADE ON DELETE CASCADE,
    scan_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
