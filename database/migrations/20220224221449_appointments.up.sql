CREATE TABLE appointments
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    appointment_date date,
    appointment_time varchar,
    clinic_name varchar,
    clinic_referral_number varchar,
    form_scan_id uuid  REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    specialization_id uuid  REFERENCES specializations (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    doctor_id uuid  REFERENCES doctors (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    oms boolean,
    mrt boolean,
    mrt_zone varchar,
    mrt_anesthesia varchar,
    user_id uuid  REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    child_id uuid  REFERENCES children (id) ON UPDATE CASCADE ON DELETE CASCADE
);