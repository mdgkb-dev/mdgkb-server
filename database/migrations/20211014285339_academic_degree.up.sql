alter table doctors drop column education;

alter table doctors
    add academic_degree varchar;

alter table doctors
    add academic_rank varchar;
