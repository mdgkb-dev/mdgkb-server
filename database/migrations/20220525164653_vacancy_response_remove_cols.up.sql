drop view vacancies_view;

ALTER TABLE vacancy_responses
    DROP COLUMN response_date;
ALTER TABLE vacancy_responses
    DROP COLUMN cover_letter;
ALTER TABLE vacancy_responses
    DROP COLUMN viewed;
ALTER TABLE vacancy_responses
    DROP COLUMN contact_info_id;
ALTER TABLE vacancy_responses
    DROP COLUMN user_id;

create or replace view vacancies_view as 
SELECT v.*, count(vr) as responses_count, count(fv) as new_responses_count
FROM vacancies v
         left join vacancy_responses vr on v.id = vr.vacancy_id
         left join form_values fv on fv.id = vr.form_value_id and fv.is_new = true
GROUP BY v.id