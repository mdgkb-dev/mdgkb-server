CREATE TABLE dpo_courses
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    description varchar,
    dpo_course_order int default 0,
    dpo_course_start date,
    listeners int default 0,
    hours int default 0,
    teacher_id uuid  REFERENCES teachers (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);