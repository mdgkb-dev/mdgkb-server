create table educations
(
    id                         uuid default uuid_generate_v4() not null
        constraint educations_pkey
            primary key,
    type                       varchar,
    institution                varchar,
    document                   varchar,
    qualification              varchar,
    doctor_id                  uuid
        constraint educations_doctor_id_fkey
            references doctors,
    education_speciality_id    uuid,
    education_certification_id uuid
        constraint educations_education_certification_id_fkey
            references education_certifications,
    education_accreditation_id uuid
        constraint educations_education_accreditation_id_fkey
            references education_accreditations,
    education_start            date,
    education_end              date,
    specialization             varchar
);