CREATE TABLE doctors_divisions
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    doctor_id uuid  REFERENCES doctors (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);

insert into doctors_divisions (division_id, doctor_id)
select  division_id, doctors.id
from doctors
join divisions on division_id = divisions.id
where division_id is not null

drop view educational_organization_academics_view;
drop view doctors_view;

ALTER TABLE doctors DROP COLUMN division_id;

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
         join humans h on h.id = doctors.human_id
         LEFT JOIN doctor_comments dc on doctors.id = dc.doctor_id
         LEFT JOIN regalias r on doctors.id = r.doctor_id
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
