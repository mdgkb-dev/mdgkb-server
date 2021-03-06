create table side_organizations
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    description varchar,
    contact_info_id uuid references contact_infos(id) on update cascade on delete cascade
);
