alter table education_certifications
    rename to certifications;

alter table certifications
    add employee_id uuid;

alter table certifications
    add constraint certifications_employees_id_fk
        foreign key (employee_id) references employees;

alter table education_accreditations
    rename to accreditations;

alter table accreditations
    add employee_id uuid;

alter table accreditations
    add constraint accreditations_employees_id_fk
        foreign key (employee_id) references employees;

