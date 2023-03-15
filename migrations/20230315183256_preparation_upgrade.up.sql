alter table preparations
    add oms bool;

alter table preparations
    add dms bool;

alter table preparations
    add laboratory bool;

alter table preparations_rules
    add item_order int;

alter table preparations_rules_groups
    add item_order int;

