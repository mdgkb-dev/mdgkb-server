alter table daily_menu_order_items
add column price numeric;

alter table daily_menu_orders
drop column price;