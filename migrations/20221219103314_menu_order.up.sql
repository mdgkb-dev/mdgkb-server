CREATE TABLE daily_menu_orders (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  item_date date,
  box_number numeric,
  price numeric,
  number numeric,
  form_value_id uuid  REFERENCES form_values (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);


CREATE TABLE daily_menu_order_items (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  item_date date,
  quantity numeric,
  daily_menu_order_id uuid  REFERENCES daily_menu_orders (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
  daily_menu_item_id uuid  REFERENCES daily_menu_items (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);
