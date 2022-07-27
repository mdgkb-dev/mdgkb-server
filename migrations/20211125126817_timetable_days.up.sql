CREATE TABLE timetable_days (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    timetable_id uuid REFERENCES timetables(id) ON UPDATE CASCADE ON DELETE CASCADE,
    timetable_pattern_id uuid REFERENCES timetable_patterns(id) ON UPDATE CASCADE ON DELETE CASCADE,
    weekday_id uuid REFERENCES weekdays(id) ON UPDATE CASCADE ON DELETE CASCADE,
    start_time time not null,
    end_time time not null,
    breaks_exists boolean,
    is_weekend boolean,
    around_the_clock boolean
);