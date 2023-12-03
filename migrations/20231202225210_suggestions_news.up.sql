drop view news_view;

create or replace view news_view as
SELECT n.*, count(nv) as views_count, count(ntt) as tag_count
FROM news n
         LEFT JOIN news_views nv ON n.id = nv.news_id
         LEFT JOIN news_to_tags ntt ON n.id = ntt.news_id
GROUP BY n.id;

