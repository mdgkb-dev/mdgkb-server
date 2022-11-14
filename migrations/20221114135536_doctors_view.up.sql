drop view educational_organization_academics_view;
drop view educational_managers_view;
drop view doctors_view;
drop view teachers_view;


create or replace view doctors_view as
SELECT
    doctors.*,
    h.name,
    h.surname,
    h.patronymic,
    h.slug,
    h.date_birth,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name,
    count(dc.id) as comments_count,
    count(r.id) as regalias_count,
    p.name as position
FROM doctors
         join employees e on e.id = doctors.employee_id
         join humans h on h.id = e.human_id
         LEFT JOIN doctor_comments dc on doctors.id = dc.doctor_id
         LEFT JOIN regalias r on doctors.id = r.employee_id
         LEFT JOIN positions p on p.id = doctors.position_id
group by dc.doctor_id, doctors.id, h.id, p.name
;

create
or replace view educational_organization_academics_view as
SELECT
    dv.full_name,
    ea.*
FROM
    educational_organization_academics ea
    JOIN doctors_view dv ON dv.id = ea.doctor_id
group by
    ea.id, dv.full_name;


create or replace view educational_managers_view as
SELECT
    educational_managers.*,
    h.name,
    h.surname,
    h.patronymic,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM educational_managers
         join employees e on e.id = educational_managers.employee_id
         join humans h on h.id = e.human_id
;

create or replace view teachers_view as
SELECT
    teachers.*,
    h.name,
    h.surname,
    h.patronymic,
    h.slug,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM teachers
    join employees e on e.id = teachers.employee_id
    join humans h on h.id = e.human_id
;
