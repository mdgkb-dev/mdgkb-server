ALTER table news_views
add column city varchar;
ALTER table news_views
add column country varchar;

ALTER table news_views
add column created_at timestamp;


SELECT * from news_views;
