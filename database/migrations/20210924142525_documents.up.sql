create table documents
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    file_info_id uuid references file_infos(id)
);
