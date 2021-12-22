CREATE TABLE partner_types
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    slug varchar,
    show_image boolean
);
