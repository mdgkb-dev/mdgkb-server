CREATE TABLE form_values (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    created_at timestamp,
    is_new boolean,
    user_id uuid REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    form_status_id uuid REFERENCES form_statuses (id) ON UPDATE CASCADE ON DELETE CASCADE
);