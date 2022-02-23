CREATE TABLE dpo_base_courses
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    description varchar,
    dpo_base_course_order int default 0,
    dpo_base_course_start date,
    cost int,
    listeners int default 0,
    hours int default 0,
    teacher_id uuid  REFERENCES teachers (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);