--drop view postgraduate_applications_view;

create
or replace view postgraduate_applications_view as
SELECT
    pa.*,
    fv.created_at,
    fv.form_status_id,
    pc.name as course_name,
    u.email,
    CONCAT_WS(' ' :: TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM
    postgraduate_applications pa
    join form_values fv on fv.id = pa.form_value_id
    join postgraduate_courses_view pc on pc.id = pa.postgraduate_course_id
    join users u on u.id = fv.user_id
    join humans h on h.id = u.human_id;