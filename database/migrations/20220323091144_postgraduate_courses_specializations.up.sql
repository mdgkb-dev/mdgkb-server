CREATE TABLE postgraduate_courses_specializations
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    main boolean,
    specialization_id uuid  REFERENCES specializations (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    postgraduate_course_id uuid  REFERENCES postgraduate_courses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
