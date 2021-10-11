CREATE TABLE vacancies (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    vacancy_date date,
    title varchar,
    description varchar,
    specialization varchar,
    salary varchar
);