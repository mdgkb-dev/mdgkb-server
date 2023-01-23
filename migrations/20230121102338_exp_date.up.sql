alter table experiences
    rename column experience_start to item_start;

alter table experiences
    alter column item_start type date using to_date(item_start::text, 'YYYY');

alter table experiences
    rename column experience_end to "item_end";

alter table experiences
    alter column item_end type date using to_date(item_end::text, 'YYYY');