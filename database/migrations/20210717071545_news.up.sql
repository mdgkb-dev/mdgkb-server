CREATE TABLE file_infos
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    filename_disk VARCHAR,
    original_name varchar
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
    file_info_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
