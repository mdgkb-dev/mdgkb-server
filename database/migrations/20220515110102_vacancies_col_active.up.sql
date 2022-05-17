ALTER TABLE vacancies
    RENAME COLUMN archived TO active;

drop view vacancies_view;

create or replace view vacancies_view as
SELECT v.*, count(vr) as responses_count, count(vrn) as new_responses_count
FROM vacancies v
         left join vacancy_responses vr on v.id = vr.vacancy_id
         left join vacancy_responses vrn on v.id = vrn.vacancy_id and vrn.viewed = true
GROUP BY v.id