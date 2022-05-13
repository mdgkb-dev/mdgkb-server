-- drop view dpo_applications_view;

create or replace view dpo_applications_view as
SELECT
    da.*,
    dc.is_nmo,
       fv.created_at
FROM dpo_applications da
    join form_values fv on fv.id = da.form_value_id
         join dpo_courses dc on dc.id = da.dpo_course_id
;