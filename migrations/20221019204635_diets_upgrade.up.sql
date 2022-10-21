CREATE TABLE diets_groups (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name varchar
);

alter table diets
    drop column timetable_id;

ALTER TABLE diets
    ADD COLUMN diet_group_id uuid  REFERENCES diets_groups(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

alter table age_periods
    rename to diet_ages;

ALTER TABLE diet_ages
    ADD COLUMN timetable_id uuid  REFERENCES timetables(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

ALTER TABLE diet_ages
    ADD COLUMN diet_id uuid  REFERENCES diets(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;