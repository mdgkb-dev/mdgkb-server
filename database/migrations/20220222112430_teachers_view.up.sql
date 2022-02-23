create or replace view teachers_view as
SELECT
    teachers.*,
    h.name,
    h.surname,
    h.patronymic,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM teachers
         join doctors d on d.id = teachers.doctor_id
         join humans h on h.id = d.human_id
;
