CREATE TABLE timetable_days
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    timetable_id uuid  REFERENCES timetables(id) ON UPDATE CASCADE ON DELETE CASCADE,
    weekday_id uuid  REFERENCES weekdays(id) ON UPDATE CASCADE ON DELETE CASCADE,
    start_time time not null ,
    end_time time,
    break_exist boolean,
    is_weekend boolean,
    break_start_time time,
    break_end_time time,
    is_custom boolean,
    custom_name varchar
);
