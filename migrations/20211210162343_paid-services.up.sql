CREATE TABLE paid_services (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    code varchar,
    nomenclature varchar,
    type varchar,
    price int
);