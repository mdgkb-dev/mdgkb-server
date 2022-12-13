alter table appointments
    drop column clinic_name;

alter table appointments
    drop column clinic_referral_number;

alter table appointments
    drop column form_scan_id;

alter table appointments
    drop column oms;

alter table appointments
    drop column mrt;

alter table appointments
    drop column mrt_zone;

alter table appointments
    drop column mrt_anesthesia;

alter table appointments
    drop column user_id;

alter table appointments
    drop column child_id;

ALTER TABLE appointments
    ADD COLUMN form_value_id uuid REFERENCES form_values(id) ON UPDATE CASCADE ON DELETE CASCADE;

CREATE TABLE appointments_types (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  oms bool,
  name varchar,
  description varchar,
  item_order integer default 0,
  form_pattern_id uuid REFERENCES form_patterns(id) ON UPDATE CASCADE ON DELETE CASCADE
);
