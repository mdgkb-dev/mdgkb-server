ALTER TABLE form_statuses
    ADD COLUMN form_status_group_id uuid  REFERENCES form_status_groups (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
