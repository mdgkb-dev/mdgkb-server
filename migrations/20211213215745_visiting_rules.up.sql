CREATE TABLE visiting_rules (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    rule_order int not null default 0,
    text VARCHAR,
    is_list_item boolean default true,
    division_id uuid REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);