create or replace view divisions_view as
SELECT
    divisions.*,
    count(dc.id) as comments_count
FROM divisions
         LEFT JOIN division_comments dc on divisions.id = dc.division_id
group by dc.division_id, divisions.id
;

