CREATE TABLE residency_courses_teachers
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    main boolean,
    teacher_id uuid  REFERENCES teachers (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    residency_course_id uuid  REFERENCES residency_courses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
