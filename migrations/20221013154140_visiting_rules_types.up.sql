CREATE TABLE visiting_rules_groups
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    color varchar,
    visiting_rule_group_order int,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);

alter table visiting_rules drop column division_id;
alter table visiting_rules drop column visiting_rule_group_id uuid REFERENCES visiting_rules_groups (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;