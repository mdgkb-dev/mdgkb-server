CREATE TABLE dpo_courses_dates
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    dpo_course_start date,
    dpo_course_end date,
    dpo_course_id uuid  REFERENCES dpo_courses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);