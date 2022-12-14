drop view if exists appointments_view;

create or replace view appointments_view as
SELECT
    a.*,
    apt.name, apt.oms,
    u.email,
    fv.is_new,
    fv.created_at,
    CONCAT_WS(' '::TEXT, humans.surname, humans.name, humans.patronymic) AS full_name
FROM appointments a
         join form_values fv on fv.id = a.form_value_id
         join users u on u.id = fv.user_id
         join humans on humans.id = u.human_id
         LEFT JOIN appointments_types apt on apt.id = a.appointment_type_id;

