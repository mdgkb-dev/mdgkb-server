CREATE TABLE vacancy_requirements (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    vacancy_id uuid  REFERENCES vacancies (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);