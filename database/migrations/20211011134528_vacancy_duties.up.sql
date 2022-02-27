CREATE TABLE vacancy_duties (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    vacancy_id uuid  REFERENCES vacancies (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);