alter table appointments
    rename column appointment_date to item_date;

alter table appointments
    rename column appointment_time to item_time;

alter table appointments
    ADD COLUMN appointment_type_id uuid  REFERENCES appointments_types(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
