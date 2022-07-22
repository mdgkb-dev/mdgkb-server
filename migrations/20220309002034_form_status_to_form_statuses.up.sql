CREATE TABLE form_status_to_form_statuses (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    form_status_id uuid REFERENCES form_statuses (id) ON UPDATE CASCADE ON DELETE CASCADE,
    child_form_status_id uuid REFERENCES form_statuses (id) ON UPDATE CASCADE ON DELETE CASCADE
);