CREATE TABLE postgraduate_applications (
id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
created_at timestamp,
is_new boolean,
 user_id uuid REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
 postgraduate_course_id uuid REFERENCES postgraduate_courses (id) ON UPDATE CASCADE ON DELETE CASCADE
);