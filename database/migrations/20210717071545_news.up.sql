CREATE TABLE preview_thumbnail_files
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    filename_disk VARCHAR
);


CREATE TABLE news
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    status VARCHAR,
    title          VARCHAR,
    preview_text   VARCHAR,
    content        text,
    slug           VARCHAR,
    published_on   date,
    description    VARCHAR,
    preview_thumbnail_file_id uuid  REFERENCES preview_thumbnail_files(id) ON UPDATE CASCADE ON DELETE CASCADE
);
