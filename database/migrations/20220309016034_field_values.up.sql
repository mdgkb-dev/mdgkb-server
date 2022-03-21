CREATE TABLE field_values
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    value_string varchar,
    value_number integer,
    value_date date,
    field_id uuid not null references fields(id),
    event_application_id uuid references event_applications(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    dpo_application_id uuid references dpo_applications(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    file_id uuid  REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
