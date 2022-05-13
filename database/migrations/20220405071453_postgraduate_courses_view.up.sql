create or replace view postgraduate_courses_view as
SELECT s.name, s.code, s.slug, pc.*
FROM postgraduate_courses pc
         JOIN postgraduate_courses_specializations pcs ON pcs.postgraduate_course_id = pc.id and pcs.main = true
         join specializations s on pcs.specialization_id = s.id
group by pc.id, s.name, s.code, s.slug;

