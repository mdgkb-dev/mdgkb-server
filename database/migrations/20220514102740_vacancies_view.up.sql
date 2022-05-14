create or replace view vacancies_view as
SELECT v.*, count(vr) as responses_count
FROM vacancies v
         left join vacancy_responses vr on v.id = vr.vacancy_id
GROUP BY v.id