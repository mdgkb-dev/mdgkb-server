CREATE TABLE education_years
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    year date,
    active boolean
);
