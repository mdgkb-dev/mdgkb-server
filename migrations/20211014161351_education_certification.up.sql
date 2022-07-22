CREATE TABLE education_certifications (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    specialization  varchar,
    certification_date varchar,
    end_date          varchar,
    place                       varchar,
    document           varchar
);