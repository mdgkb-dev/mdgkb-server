create or replace view educational_managers_view as
SELECT
    educational_managers.*,
    h.name,
    h.surname,
    h.patronymic,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM educational_managers
         join doctors d on d.id = educational_managers.doctor_id
         join humans h on h.id = d.human_id
;
