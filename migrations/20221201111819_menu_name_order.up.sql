ALTER TABLE daily_menus
    ADD COLUMN item_order integer;

ALTER TABLE daily_menus
    ADD COLUMN name varchar;

ALTER TABLE daily_menus
    ADD COLUMN active boolean;

alter table daily_menus
    rename column daily_menu_date to item_date;

