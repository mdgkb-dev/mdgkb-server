create or replace view dpo_courses_view as
SELECT min(dcd.dpo_course_start) as min_dpo_course_start , min(dpo_course_end) as min_dpo_course_end, dc.*
FROM dpo_courses dc
         JOIN dpo_courses_dates dcd ON dcd.dpo_course_id = dc.id
group by dc.id;

