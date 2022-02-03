CREATE TABLE heads
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    human_id uuid  REFERENCES humans (id) ON UPDATE CASCADE ON DELETE CASCADE,
    position          VARCHAR,
    tags              VARCHAR,
    photo_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    timetable_id uuid  REFERENCES timetables(id) ON UPDATE CASCADE ON DELETE CASCADE,
    contact_info_id uuid references contact_infos(id) on update cascade on delete cascade,
    academic_degree varchar,
    academic_rank varchar,
    is_main boolean
);
