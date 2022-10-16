CREATE TABLE residency_course_practice_places (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name varchar,
  link varchar,
  residency_course_practice_place_order int,
  residency_course_id uuid REFERENCES residency_courses(id) ON UPDATE CASCADE ON DELETE CASCADE
);