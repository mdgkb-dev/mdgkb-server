create or replace view residency_courses_view as
SELECT s.name, s.code, s.slug, pc.*
FROM residency_courses pc
         JOIN residency_courses_specializations pcs ON pcs.residency_course_id = pc.id and pcs.main = true
         join specializations s on pcs.specialization_id = s.id
group by pc.id, s.name, s.code, s.slug;

