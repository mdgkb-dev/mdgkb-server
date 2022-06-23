CREATE TABLE points_achievements
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    points INT
);


CREATE TABLE residency_applications_points_achievement
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    residency_application_id  uuid  REFERENCES residency_applications(id) ON UPDATE CASCADE ON DELETE CASCADE   t,
    points_achievement_id  uuid  REFERENCES points_achievements(id) ON UPDATE CASCADE ON DELETE CASCADE,
    file_info_id uuid  REFERENCES file_infos (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    approved boolean
);
