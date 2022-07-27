create table document_types
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    description varchar,
    public_document_type_id uuid REFERENCES public_document_types (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL

    -- single_scan boolean,
    -- scan_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
