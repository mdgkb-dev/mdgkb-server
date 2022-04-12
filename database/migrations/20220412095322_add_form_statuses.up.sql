ALTER TABLE
    form_statuses
ADD
    column icon_id uuid REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;