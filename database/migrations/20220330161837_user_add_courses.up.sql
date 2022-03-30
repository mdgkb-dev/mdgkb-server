ALTER TABLE users
    ADD COLUMN dpo_application_id uuid references dpo_applications(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

ALTER TABLE users
    ADD COLUMN postgraduate_application_id uuid REFERENCES postgraduate_applications(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

ALTER TABLE users
    ADD COLUMN candidate_application_id uuid references candidate_applications(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
