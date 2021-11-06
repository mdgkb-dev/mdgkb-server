create table document_field_values
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_caadab631597b6ad85f1d61f08c"
        primary key,
    value_string varchar,
    value_number integer,
    value_date date,
    document_id uuid not null references documents,
    document_type_field_id uuid not null references document_type_fields
);
