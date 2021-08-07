CREATE TABLE humen
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    surname VARCHAR,
    patronymic VARCHAR,
    is_male boolean,
    date_birth   date
);
