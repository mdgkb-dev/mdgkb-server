CREATE TABLE schedule_items
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    description varchar,
    schedule_id uuid  REFERENCES schedules(id) ON UPDATE CASCADE ON DELETE CASCADE,
    start_time time not null ,
    end_time time
);
