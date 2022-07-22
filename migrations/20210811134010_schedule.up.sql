CREATE TABLE schedules
(
    id         uuid DEFAULT uuid_generate_v4()                               NOT NULL PRIMARY KEY,
    description varchar,
    name varchar
);
