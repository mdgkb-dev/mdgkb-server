drop view if exists appointments_view;

create or replace view appointments_view as
SELECT
    a.*,
    at.name, at.description, at.oms, at.item_order,
    u.email,
    fv.is_new,
    fv.created_at,
    CONCAT_WS(' '::TEXT, humans.surname, humans.name, humans.patronymic) AS full_name
FROM appointments a
         join form_values fv on fv.id = a.form_value_id
         join users u on u.id = fv.user_id
         join humans on humans.id = u.human_id
         join appointments_types at on a.appointment_type_id = at.id;