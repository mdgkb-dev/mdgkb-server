create table vacancies_responses_to_documents
(
    id uuid default uuid_generate_v4() not null,
    document_id uuid references documents on delete cascade,
    vacancy_response_id uuid references vacancy_responses on delete cascade
);

