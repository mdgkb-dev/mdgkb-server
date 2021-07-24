CREATE TABLE side_organizations
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    site varchar,
    phone varchar,
    address varchar
);

