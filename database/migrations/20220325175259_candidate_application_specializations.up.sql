CREATE TABLE candidate_application_specializations
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    main boolean,
    candidate_application_id uuid  REFERENCES candidate_applications (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    specialization_id uuid  REFERENCES specializations (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
