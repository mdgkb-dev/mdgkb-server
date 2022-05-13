CREATE TABLE dpo_courses_teachers
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    main boolean,
    teacher_id uuid  REFERENCES teachers (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    dpo_course_id uuid  REFERENCES dpo_courses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);