CREATE TABLE residency_courses
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    description varchar,
    years integer,
    listeners integer,
    education_form varchar
);
