CREATE TABLE carousel_slides
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    title        varchar,
    content        text,
    carousel_id uuid  REFERENCES carousels(id) ON UPDATE CASCADE ON DELETE CASCADE,
    file_info_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
