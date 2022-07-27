CREATE TABLE visits (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    date timestamp,
    entered boolean default false,
    exited boolean default false,
    application_car_id uuid REFERENCES applications_cars (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);