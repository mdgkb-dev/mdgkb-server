drop view dpo_applications_view;

create or replace view dpo_applications_view as
SELECT
    da.*,
    dc.is_nmo
FROM dpo_applications da
         join dpo_courses dc on dc.id = da.dpo_course_id;