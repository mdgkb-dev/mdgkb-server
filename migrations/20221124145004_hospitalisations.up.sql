drop table hospitalizations;

CREATE TABLE hospitalizations_types (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  paid bool,
  policy_type varchar,
  treatment_type varchar,
  stay_type varchar,
  referral_type varchar
);

CREATE TABLE hospitalizations (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name varchar,
  hospitalization_date date,
  form_value_id uuid REFERENCES form_values(id) ON UPDATE CASCADE ON DELETE CASCADE,
  division_id uuid REFERENCES divisions(id) ON UPDATE CASCADE ON DELETE CASCADE,
  hospitalization_type_id uuid REFERENCES hospitalizations_types(id) ON UPDATE CASCADE ON DELETE CASCADE,
  medical_scan_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);