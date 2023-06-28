CREATE TABLE diplomas (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    university_name VARCHAR,
    university_end_date timestamp,
    number VARCHAR,
    series VARCHAR,
    date timestamp,
    speciality_code VARCHAR,
    speciality VARCHAR
);