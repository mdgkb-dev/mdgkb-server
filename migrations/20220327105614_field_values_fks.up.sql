ALTER TABLE field_values
    ADD COLUMN postgraduate_application_id uuid REFERENCES postgraduate_applications(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE field_values
    ADD COLUMN candidate_application_id uuid references candidate_applications(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

