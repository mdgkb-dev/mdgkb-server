drop view doctors_view;

create or replace view doctors_view as
SELECT
    doctors.*,
    h.name,
    h.surname,
    h.patronymic,
       h.slug,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM doctors
         join humans h on h.id = doctors.human_id
;
