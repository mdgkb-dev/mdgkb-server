create table documents
(
    id uuid default uuid_generate_v4() not null primary key,
    document_type_id uuid REFERENCES document_types (id)
);
