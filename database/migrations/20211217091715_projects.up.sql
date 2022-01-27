CREATE TABLE projects
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    title varchar,
    content text,
    image_id uuid REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    slug VARCHAR
);
