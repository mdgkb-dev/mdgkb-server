ALTER TABLE doctors
    ADD COLUMN timetable_id uuid  REFERENCES timetables(id) ON UPDATE CASCADE ON DELETE CASCADE;