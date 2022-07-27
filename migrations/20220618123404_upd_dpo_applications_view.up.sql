drop view dpo_applications_view;

create
or replace view dpo_applications_view as
SELECT
    da.*,
    dc.is_nmo,
    fv.created_at,
    dc.name as course_name,
    fv.form_status_id,
    u.email,
    CONCAT_WS(' ' :: TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM
    dpo_applications da
    join form_values fv on fv.id = da.form_value_id
    join dpo_courses dc on dc.id = da.dpo_course_id
    join users u on u.id = fv.user_id
    join humans h on h.id = u.human_id;