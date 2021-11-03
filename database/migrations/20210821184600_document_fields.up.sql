create table document_fields
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_504d1a17f1681be11d94673ba31"
        primary key,
    name varchar not null,
    value_type_id uuid
        constraint "FK_f021d1969d6682de1f265c62a6c"
            references value_types,
    document_field_order integer,
    document_id uuid
        constraint "FK_040777158438fdb7a2ca0d9a3bd"
        references documents
        on delete cascade
);

