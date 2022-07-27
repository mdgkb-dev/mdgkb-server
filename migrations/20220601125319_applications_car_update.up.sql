ALTER TABLE applications_cars
    DROP COLUMN car_number;
ALTER TABLE applications_cars
    DROP COLUMN car_brand;
ALTER TABLE applications_cars
    DROP COLUMN hospitalization_date;
ALTER TABLE applications_cars
    DROP COLUMN user_id;
ALTER TABLE applications_cars
    DROP COLUMN moved_in;
ALTER TABLE applications_cars
    DROP COLUMN moved_out;

ALTER TABLE applications_cars
    ADD COLUMN form_value_id uuid REFERENCES form_values (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
