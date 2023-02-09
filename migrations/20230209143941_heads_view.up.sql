create or replace view heads_view as
SELECT
    heads.*,
    h.name,
    h.surname,
    h.patronymic,
    h.slug,
    h.date_birth,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name,
    h.is_male
FROM heads
         join employees e on e.id = heads.employee_id
         join humans h on h.id = e.human_id;