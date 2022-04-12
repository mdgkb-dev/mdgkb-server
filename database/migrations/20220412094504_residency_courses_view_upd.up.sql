drop view residency_courses_view;

create or replace view residency_courses_view as
SELECT s.name, s.code, s.slug, sy.id as start_year_id, ey.id as end_year_id, sy.year as start_year, ey.year as end_year, rc.*
FROM residency_courses rc
         JOIN residency_courses_specializations rcs ON rcs.residency_course_id = rc.id and rcs.main = true
         join specializations s on rcs.specialization_id = s.id
         left join education_years sy on rc.start_year_id = sy.id
         left join education_years ey on rc.end_year_id = ey.id
group by rc.id, s.name, s.code, s.slug, sy.year, s.code, s.slug, s.name, ey.year;



