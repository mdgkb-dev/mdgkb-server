CREATE TABLE page_images (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    description varchar,
    page_id uuid REFERENCES pages (id) ON UPDATE CASCADE ON DELETE CASCADE,
    file_info_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
