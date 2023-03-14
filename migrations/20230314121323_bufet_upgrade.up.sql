alter table daily_menu_items
    add cook boolean;

alter table daily_menu_items
    add tomorrow_available bool;

alter table daily_menu_items
    add proteins int;
alter table daily_menu_items
    add fats int;
alter table daily_menu_items
    add carbohydrates int;
alter table daily_menu_items
    add dietary boolean;
alter table daily_menu_items
    add lean boolean;

alter table dishes_samples
    add proteins int;
alter table dishes_samples
    add fats int;
alter table dishes_samples
    add carbohydrates int;
alter table dishes_samples
    add dietary boolean;
alter table dishes_samples
    add lean boolean;

alter table daily_menus
    add end_time time;


