CREATE TABLE preparations_rules_groups
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    preparation_id uuid  REFERENCES preparations (id) ON UPDATE CASCADE ON DELETE CASCADE
);
