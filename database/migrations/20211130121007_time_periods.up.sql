CREATE TABLE time_periods (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    timetable_day_id uuid REFERENCES timetable_days(id) ON UPDATE CASCADE ON DELETE CASCADE,
    start_time time not null,
    end_time time not null
);