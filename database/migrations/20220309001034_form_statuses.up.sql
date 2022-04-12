CREATE TABLE form_statuses (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    label VARCHAR,
    color VARCHAR,
    mod_action_name VARCHAR,
    user_action_name VARCHAR,
    is_editable boolean,
    icon_id uuid  REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);