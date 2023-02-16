drop view dpo_courses_view;

alter table dpo_courses rename to nmo_courses;
alter table dpo_courses_teachers rename to nmo_courses_teachers;
alter table dpo_courses_specializations rename to nmo_courses_specializations;

alter table nmo_courses drop column is_nmo;
alter table nmo_courses rename column link_nmo to link;

alter table nmo_courses
    add main_teacher_id uuid;

alter table nmo_courses
    add constraint npo_courses_employees_id_fk
        foreign key (main_teacher_id) references employees
            on update cascade on delete cascade;

alter table nmo_courses_teachers rename column dpo_course_id to nmo_course_id;
alter table nmo_courses_specializations rename column dpo_course_id to nmo_course_id;
drop table dpo_document_types;