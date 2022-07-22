CREATE TABLE vacancies (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    vacancy_date date,
    title varchar,
    slug varchar,
    specialization varchar,
    archived boolean,
    min_salary int,
    max_salary int,
    salary_comment varchar,
    requirements varchar,
    experience varchar,
    schedule varchar,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    contact_doctor_id uuid  REFERENCES doctors (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    contact_info_id uuid  REFERENCES contact_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);