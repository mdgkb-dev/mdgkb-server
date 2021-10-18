CREATE TABLE vacancies (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    vacancy_date date,
    title varchar,
    description varchar,
    specialization varchar,
    archived boolean,
    salary varchar,
    requirements varchar,
    experience varchar,
    duties varchar,
    schedule varchar,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);