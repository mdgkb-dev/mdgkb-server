create sequence daily_menu_orders_number_seq
    as integer;

alter table daily_menu_orders
    alter column number type integer using number::integer;

alter table daily_menu_orders
    alter column number set not null;

alter table daily_menu_orders
    alter column number set default nextval('public.daily_menu_orders_number_seq'::regclass);

alter sequence daily_menu_orders_number_seq owned by daily_menu_orders.number;

create unique index daily_menu_orders_number_uindex
    on daily_menu_orders (number);

