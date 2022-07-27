CREATE TABLE candidate_applications (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    candidate_exam_id uuid REFERENCES candidate_exams (id) ON UPDATE CASCADE ON DELETE CASCADE,
    form_value_id uuid  REFERENCES form_values (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);