alter table daily_menus
    add start_time time;

alter table daily_menu_items
    add from_other_menu boolean;

