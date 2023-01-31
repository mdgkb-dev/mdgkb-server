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
    p.name as position,
    h.is_male
FROM doctors
         join employees e on e.id = doctors.employee_id
         join humans h on h.id = e.human_id
         LEFT JOIN doctor_comments dc on doctors.id = dc.doctor_id
         LEFT JOIN regalias r on doctors.id = r.employee_id
         LEFT JOIN positions p on p.id = doctors.position_id
group by dc.doctor_id, doctors.id, h.id, p.name
;

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



create or replace view teachers_view as
SELECT
    teachers.*,
    h.name,
    h.surname,
    h.patronymic,
    h.slug,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name,
    h.is_male,
    h.date_birth
FROM teachers
         join employees e on e.id = teachers.employee_id
         join humans h on h.id = e.human_id;
;
