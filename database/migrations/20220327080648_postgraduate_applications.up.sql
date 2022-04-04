CREATE TABLE postgraduate_applications (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    postgraduate_course_id uuid REFERENCES postgraduate_courses (id) ON UPDATE CASCADE ON DELETE CASCADE,
    form_value_id uuid  REFERENCES form_values (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);