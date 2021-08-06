CREATE TABLE news_images (
                               id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
                               description varchar,
                               news_id uuid  REFERENCES news (id) ON UPDATE CASCADE ON DELETE CASCADE,
                               file_info_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
