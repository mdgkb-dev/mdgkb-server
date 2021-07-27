create table normative_document_types
(
   id uuid default uuid_generate_v4() not null primary key,
   name varchar
);

create table normative_documents
(
   id uuid default uuid_generate_v4() not null primary key,
   normative_document_type_id uuid references normative_document_types(id) on update cascade on delete cascade,
   name varchar,
   file_info_id uuid references file_infos(id)
);
