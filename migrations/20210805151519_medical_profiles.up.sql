create table medical_profiles
(
    id   uuid default uuid_generate_v4() not null constraint medical_profile_pk primary key,
    name varchar,
    description varchar,
    icon_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);