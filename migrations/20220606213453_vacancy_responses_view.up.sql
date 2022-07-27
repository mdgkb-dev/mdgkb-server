-- drop view vacancy_responses_view;

create
or replace view vacancy_responses_view as
SELECT
    vr.*,
    fv.created_at,
    v.title
FROM
    vacancy_responses vr
    join form_values fv on fv.id = vr.form_value_id
    join vacancies v on v.id = vr.vacancy_id;