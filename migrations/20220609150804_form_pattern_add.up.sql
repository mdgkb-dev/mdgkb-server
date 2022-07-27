ALTER TABLE form_patterns
    ADD COLUMN personal_data_agreement_id uuid  REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
ALTER TABLE form_patterns
    ADD COLUMN with_personal_data_agreement boolean default false;
