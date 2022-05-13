create table hospitalizations_to_document_types
(
    id uuid default uuid_generate_v4() not null,
    document_type_id uuid references document_types on delete cascade,
    hospitalization_id uuid references hospitalizations on delete cascade
);

