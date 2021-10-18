CREATE TABLE educations (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    type varchar,
    institution varchar,
    document varchar,
    qualification varchar,
    doctor_id uuid REFERENCES doctors (id),
    education_speciality_id uuid,
    education_certification_id uuid REFERENCES education_certifications (id),
    education_accreditation_id uuid REFERENCES education_accreditations (id)
);