CREATE TABLE residency_applications (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    residency_course_id uuid REFERENCES residency_courses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    form_value_id uuid  REFERENCES form_values (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);