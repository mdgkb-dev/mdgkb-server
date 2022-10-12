drop view divisions_view;

ALTER TABLE divisions ADD COLUMN is_center boolean;

create view divisions_view
as
SELECT divisions.*,
       b.name as building_id,
       count(dc.id) AS comments_count,
       tn.number    AS phone,
       em.address    AS email
FROM divisions
         LEFT JOIN division_comments dc ON divisions.id = dc.division_id
         LEFT JOIN contact_infos ci ON ci.id = divisions.contact_info_id
         LEFT JOIN entrances e ON e.id = divisions.entrance_id
         LEFT JOIN buildings b ON e.building_id = b.id
         LEFT JOIN telephone_numbers tn ON tn.contact_info_id = ci.id AND tn.main = true
         LEFT JOIN emails em ON em.contact_info_id = ci.id AND em.main = true
GROUP BY dc.division_id, divisions.id, tn.number, em.address, b.name;

insert into divisions (name, info, address, is_center)
select name, info, address, true from centers;

drop table centers;