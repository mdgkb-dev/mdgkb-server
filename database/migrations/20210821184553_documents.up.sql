create table documents
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_2e1aa55eac1947ddf3221506edb"
        primary key,
    name varchar not null,
    file_info_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
