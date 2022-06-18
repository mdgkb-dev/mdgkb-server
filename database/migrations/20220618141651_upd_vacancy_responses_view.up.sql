drop view vacancy_responses_view;

create
or replace view vacancy_responses_view as
SELECT
    vr.*,
    fv.created_at,
    fv.form_status_id,
    v.title,
    u.email,
    CONCAT_WS(' ' :: TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM
    vacancy_responses vr
    join form_values fv on fv.id = vr.form_value_id
    join vacancies v on v.id = vr.vacancy_id
    join users u on u.id = fv.user_id
    join humans h on h.id = u.human_id;