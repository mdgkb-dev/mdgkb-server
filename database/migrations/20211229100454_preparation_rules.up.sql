CREATE TABLE preparations_rules
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    rule_time varchar,
    preparation_rules_group_id uuid  REFERENCES preparations_rules_groups (id) ON UPDATE CASCADE ON DELETE CASCADE
);
