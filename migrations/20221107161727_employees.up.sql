CREATE TABLE employees
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    uuid uuid DEFAULT uuid_generate_v4() not null ,
    human_id uuid  REFERENCES humans (id) ON UPDATE CASCADE ON DELETE CASCADE,
    academic_degree varchar,
    academic_rank varchar
);

ALTER TABLE doctors ADD COLUMN employee_id uuid  REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
ALTER TABLE regalias ADD COLUMN employee_id uuid  REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
ALTER TABLE educations ADD COLUMN employee_id uuid  REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
ALTER TABLE experiences ADD COLUMN employee_id uuid  REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
ALTER TABLE certificates ADD COLUMN employee_id uuid  REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

ALTER TABLE educational_managers ADD COLUMN employee_id uuid  REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
ALTER TABLE teachers ADD COLUMN employee_id uuid  REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
ALTER TABLE heads ADD COLUMN employee_id uuid  REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

insert into employees (human_id, academic_degree, academic_rank)
select human_id, academic_degree, academic_rank from doctors;

insert into employees (human_id, academic_degree, academic_rank)
select human_id, academic_degree, academic_rank from heads;

update doctors d set employee_id =
employees.id from employees
where employees.human_id = d.human_id;

update heads h set employee_id =
employees.id from employees
where employees.human_id = h.human_id;

update regalias r set employee_id =
doctors.employee_id from doctors
where doctors.id = r.doctor_id;

update educations e set employee_id =
doctors.employee_id from doctors
where doctors.id = e.doctor_id;

update experiences e set employee_id =
doctors.employee_id from doctors
where doctors.id = e.doctor_id;

update certificates c set employee_id =
doctors.employee_id from doctors
where doctors.id = c.doctor_id;

update educational_managers em set employee_id =
doctors.employee_id from doctors
where doctors.id = em.doctor_id;

update teachers t set employee_id =
doctors.employee_id from doctors
where doctors.id = t.doctor_id;

drop view educational_organization_academics_view;
drop view educational_managers_view;
drop view doctors_view;
drop view teachers_view;

alter table doctors drop column human_id;
alter table doctors drop column academic_rank;
alter table doctors drop column academic_degree;

alter table heads drop column human_id;
alter table heads drop column academic_rank;
alter table heads drop column academic_degree;

alter table regalias drop column doctor_id;
alter table regalias drop column head_id;

create or replace view doctors_view as
SELECT
    doctors.*,
    h.name,
    h.surname,
    h.patronymic,
    h.slug,
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

alter table educations drop column doctor_id;
alter table experiences drop column doctor_id;
alter table certificates drop column doctor_id;

create or replace view employees_view as
SELECT
    e.*,
    h.name,
    h.surname,
    h.patronymic,
    h.slug,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM employees e
         join humans h on h.id = e.human_id;
