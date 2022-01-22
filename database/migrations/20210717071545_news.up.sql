create table file_infos
(
    id uuid default uuid_generate_v4() not null primary key,
    original_name varchar,
    file_system_path varchar
);


CREATE TABLE news
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    status VARCHAR,
    title          VARCHAR,
    preview_text   VARCHAR,
    content        text,
    slug           VARCHAR,
    published_on   timestamp,
    main boolean default false,
    is_article boolean default false,
    article_link varchar,
    sub_main boolean default false,
    description    VARCHAR,
    file_info_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
