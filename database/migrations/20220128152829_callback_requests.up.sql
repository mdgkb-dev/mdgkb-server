CREATE TABLE callback_requests
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    phone varchar,
    name varchar,
    description varchar
);