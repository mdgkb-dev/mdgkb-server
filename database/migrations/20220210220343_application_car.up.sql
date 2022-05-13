CREATE TABLE applications_cars
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    car_number varchar,
    car_brand varchar,
    hospitalization_date date,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    gate_id uuid  REFERENCES gates (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    user_id uuid  REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);