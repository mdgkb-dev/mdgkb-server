ALTER TABLE divisions
    ADD COLUMN hospitalization_contact_info_id uuid REFERENCES contact_infos(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE divisions
    ADD COLUMN hospitalization_doctor_id uuid REFERENCES doctors(id) ON UPDATE CASCADE ON DELETE CASCADE;

