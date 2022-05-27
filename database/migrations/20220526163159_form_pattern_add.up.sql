ALTER TABLE form_patterns
    ADD COLUMN form_status_group_id uuid  REFERENCES form_status_groups (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
ALTER TABLE form_patterns
    ADD COLUMN default_form_status_id uuid  REFERENCES form_statuses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
