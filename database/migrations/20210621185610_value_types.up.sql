DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'value_type_value_relation_enum') THEN
    CREATE TYPE value_type_value_relation_enum AS ENUM ('simple', 'oneToMany', 'manyToMany');
END IF;
END$$;

create table value_types
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_38e761371cca3cc3051fb60674a"
        primary key,
    name varchar,
    value_relation value_type_value_relation_enum not null
);
