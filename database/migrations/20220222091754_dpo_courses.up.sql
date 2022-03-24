CREATE TABLE dpo_courses (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    description varchar,
    is_nmo boolean,
    link_nmo varchar,
    cost int,
    dpo_course_order int default 0,
    listeners int default 0,
    hours int default 0,
    specialization_id uuid  REFERENCES specializations (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    form_pattern_id uuid references form_patterns on update cascade on delete cascade
);