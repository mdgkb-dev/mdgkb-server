create table news_slide_buttons (
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    color varchar,
    background_color varchar,
    news_slide_button_order int not null default 0,
    link varchar,
    border_color varchar,
    shadow boolean,
    news_slide_id uuid  REFERENCES news_slides(id) ON UPDATE CASCADE ON DELETE CASCADE
);