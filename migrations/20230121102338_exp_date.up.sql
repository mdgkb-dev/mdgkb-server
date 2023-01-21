alter table experiences
    rename column experience_start to item_end;

alter table experiences
    alter column start type date using to_date(start::text, 'YYYY');

alter table experiences
    rename column experience_end to "item_end";

alter table experiences
    alter column "end" type date using to_date(start::text, 'YYYY');

