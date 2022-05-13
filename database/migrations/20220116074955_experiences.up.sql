CREATE TABLE experiences
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    doctor_id uuid  REFERENCES doctors (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    experience_start int,
    experience_end int,
    place varchar,
    position varchar
);