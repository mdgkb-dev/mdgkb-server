CREATE TABLE dishes_groups (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name varchar,
  dishes_group_order int
);

CREATE TABLE dishes_samples (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name varchar,
  price int,
  caloric int,
  dishes_group_id uuid REFERENCES dishes_groups(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE daily_menus (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  daily_menu_date date
);

CREATE TABLE daily_menu_item (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name varchar,
  weight int,
  price int,
  daily_menu_item_order int,
  available boolean,
  daily_menu_id uuid REFERENCES daily_menus(id) ON UPDATE CASCADE ON DELETE CASCADE
);