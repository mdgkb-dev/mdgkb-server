CREATE TABLE residency_courses
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    description varchar,
    years integer,
    free_places integer,
    paid_places integer,
    cost integer,
    education_form varchar,
    form_pattern_id uuid references form_patterns on update cascade on delete cascade,

    start_year_id  uuid  REFERENCES education_years(id) ON UPDATE CASCADE ON DELETE CASCADE,
    end_year_id  uuid  REFERENCES education_years(id) ON UPDATE CASCADE ON DELETE CASCADE,

    annotation_id  uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    program_id  uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    plan_id  uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    schedule_id  uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
