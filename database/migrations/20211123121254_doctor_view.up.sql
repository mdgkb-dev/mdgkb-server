create view doctors_view as
SELECT
       doctors.*,
       h.name,
       h.surname,
       h.patronymic,
       CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM doctors
join humans h on h.id = doctors.human_id
;
