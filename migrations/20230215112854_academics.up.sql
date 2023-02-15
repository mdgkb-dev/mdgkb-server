alter table educational_organization_academics
    add employee_id uuid;

alter table educational_organization_academics
    add constraint educational_organization_academics_employees_id_fk
        foreign key (employee_id) references employees
            on update cascade on delete cascade;

update educational_organization_academics
set employee_id = d.employee_id from
    doctors d
        join employees e on d.employee_id = e.id
where d.id = educational_organization_academics.doctor_id;

alter table educational_organization_academics
    rename to educational_academics;

alter table educational_organization_academics_view
    rename to educational_academics_view;

drop view educational_organization_academics_view;

alter table educational_academics
    drop column doctor_id;

create
    or replace view educational_academics_view as
SELECT
    h.date_birth, h.is_male,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name,
    ea.*
FROM
    educational_academics ea
        join employees e on e.id = ea.employee_id
        join humans h on e.human_id = h.id
group by
    ea.id, h.id;












