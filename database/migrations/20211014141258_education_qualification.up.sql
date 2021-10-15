CREATE TABLE education_qualifications (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    specialization  varchar,
    startDate       varchar,
    endDate         varchar,
    document        varchar
);