alter table employees
    add part_time bool default false;

drop view employees_view;
create or replace view employees_view as
SELECT
    e.*,
    h.name,
    h.surname,
    h.patronymic,
    h.slug,
    h.date_birth,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name,
    h.is_male
FROM employees e
         join humans h on h.id = e.human_id;
