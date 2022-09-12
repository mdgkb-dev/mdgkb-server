CREATE TABLE age_periods
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar
);


CREATE TABLE diets
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    short_name varchar,
    site_name varchar,
    diabetes boolean,
    mother_diet_id uuid  REFERENCES diets (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    timetable_id uuid  REFERENCES timetables (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    age_period_id uuid  REFERENCES age_periods (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);

alter table schedule_items
add column timetable_day_id uuid  REFERENCES timetables (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;


CREATE TABLE dishes
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    weight varchar,
    schedule_item_id uuid  REFERENCES schedule_items (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);