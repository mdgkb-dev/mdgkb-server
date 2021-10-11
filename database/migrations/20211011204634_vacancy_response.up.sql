CREATE TABLE vacancy_responses (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    response_date date,
    cover_letter varchar,
    vacancy_id uuid  REFERENCES vacancies (id) ON UPDATE CASCADE ON DELETE CASCADE
);