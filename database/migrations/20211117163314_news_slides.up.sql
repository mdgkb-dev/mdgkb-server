create table news_slides (
    id uuid default uuid_generate_v4() not null primary key,
    title varchar,
    content text,
    color varchar,
    news_slide_order int not null default 0,
    desktop_img_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    laptop_img_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    tablet_img_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    mobile_img_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);