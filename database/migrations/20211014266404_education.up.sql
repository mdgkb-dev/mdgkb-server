CREATE TABLE educations (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    type varchar,
    institution varchar,
    document varchar,
    doctor_id uuid REFERENCES doctors (id),
    education_speciality_id uuid REFERENCES education_specialities (id),
    education_certification_id uuid REFERENCES education_certifications (id),
    education_qualification_id uuid REFERENCES education_qualifications (id)
);