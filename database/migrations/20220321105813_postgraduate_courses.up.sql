CREATE TABLE postgraduate_courses
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    description varchar,
    years integer,
    education_form varchar,
    questions_file_id  uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
     program_file_id  uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
     calendar_id  uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    form_pattern_id uuid references form_patterns on update cascade on delete cascade,
    document_type_id uuid REFERENCES document_types (id) ON UPDATE CASCADE ON DELETE CASCADE
);
