CREATE TABLE humans
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    surname VARCHAR,
    patronymic VARCHAR,
    is_male boolean,
    date_birth   date,
    contact_info_id uuid references contact_infos(id) on update cascade on delete cascade,
    photo_id uuid references file_infos(id) on update cascade on delete cascade
);
