create or replace view news_view as
SELECT n.*, count(nv) as views_count
FROM news n
         LEFT JOIN news_views nv ON n.id = nv.news_id
GROUP BY n.id