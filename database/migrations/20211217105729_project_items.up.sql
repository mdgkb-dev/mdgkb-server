CREATE TABLE project_items
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    title varchar,
    content text,
    project_id uuid REFERENCES projects (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
