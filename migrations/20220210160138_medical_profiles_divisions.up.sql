CREATE TABLE medical_profiles_divisions
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    medical_profile_id uuid  REFERENCES medical_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);