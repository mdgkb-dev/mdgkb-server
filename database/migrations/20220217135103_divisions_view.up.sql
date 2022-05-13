create or replace view divisions_view as
SELECT divisions.id,
       divisions.name,
       divisions.show,
       divisions.info,
       divisions.address,
       divisions.slug,
       divisions.show_common_visiting_rules,
       divisions.floor_id,
       divisions.entrance_id,
       divisions.contact_info_id,
       divisions.timetable_id,
       divisions.schedule_id,
       divisions.hospitalization_contact_info_id,
       divisions.hospitalization_doctor_id,
       count(dc.id) AS comments_count,
       tn.number as phone,
       e.address as email
FROM divisions
         LEFT JOIN division_comments dc ON divisions.id = dc.division_id
         LEFT JOIN contact_infos ci ON ci.id = divisions.contact_info_id
         LEFT JOIN telephone_numbers tn ON tn.contact_info_id = ci.id and tn.main = true
         LEFT JOIN emails e ON e.contact_info_id = ci.id and e.main = true
GROUP BY dc.division_id, divisions.id, tn.number, e.address;
;

