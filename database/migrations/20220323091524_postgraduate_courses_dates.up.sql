CREATE TABLE postgraduate_courses_dates
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    postgraduate_course_start date,
    postgraduate_course_end date,
    postgraduate_course_id uuid  REFERENCES postgraduate_courses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
