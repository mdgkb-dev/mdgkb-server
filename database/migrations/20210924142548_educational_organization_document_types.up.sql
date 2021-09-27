create table educational_organization_document_types
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar
);
