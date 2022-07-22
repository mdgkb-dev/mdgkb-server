create table documents_types_for_vacancies
(
    id uuid default uuid_generate_v4() not null,
    document_type_id uuid references document_types on delete cascade
);

