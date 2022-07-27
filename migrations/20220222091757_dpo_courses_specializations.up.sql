CREATE TABLE dpo_courses_specializations
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    main boolean,
    specialization_id uuid  REFERENCES specializations (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    dpo_course_id uuid  REFERENCES dpo_courses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);