CREATE TABLE candidate_exams (
 id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
 education_year_id uuid REFERENCES education_years (id) ON UPDATE CASCADE ON DELETE CASCADE,
 form_pattern_id uuid references form_patterns on update cascade on delete cascade
);