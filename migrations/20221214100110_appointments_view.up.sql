drop view if exists hospitalizations_view;

create or replace view appointments_view as
SELECT
    h.*,
    ht.paid, ht.policy_type, ht.treatment_type, ht.stay_type, ht.referral_type,
    u.email,
    fv.is_new,
    fv.created_at,
    CONCAT_WS(' '::TEXT, humans.surname, humans.name, humans.patronymic) AS full_name
FROM hospitalizations h
         join form_values fv on fv.id = h.form_value_id
         join users u on u.id = fv.user_id
         join humans on humans.id = u.human_id
         LEFT JOIN hospitalizations_types ht on ht.id = h.hospitalization_type_id;

