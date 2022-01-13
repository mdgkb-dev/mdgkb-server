CREATE TABLE donor_rules_users
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    donor_rule_id uuid unique REFERENCES donor_rules (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    user_id uuid unique REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
        UNIQUE (donor_rule_id, user_id)

);
