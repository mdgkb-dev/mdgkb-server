CREATE TABLE education_accreditations (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    specialization  varchar,
    start_date       varchar,
    end_date         varchar,
    document        varchar
);