CREATE TABLE residency_courses_specializations
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    main boolean,
    specialization_id uuid  REFERENCES specializations (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    residency_course_id uuid  REFERENCES residency_courses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
