CREATE TABLE postgraduate_course_plans (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    plan_id  uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    year date,
    postgraduate_course_id uuid REFERENCES postgraduate_courses (id) ON UPDATE CASCADE ON DELETE CASCADE
);