CREATE TABLE division_videos
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    you_tube_video_id varchar
);