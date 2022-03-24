CREATE TABLE postgraduate_courses
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    description varchar,
    questions_file_id  uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
