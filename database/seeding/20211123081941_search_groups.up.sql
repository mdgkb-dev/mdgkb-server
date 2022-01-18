insert into public.search_groups (id, search_group_order, route, search_group_table, search_column, label, label_column, value_column, key)
values  ('4f4cf39a-1125-4989-8c11-b4335517f31d', 4, '/positions', 'positions', 'name', null, 'name', 'id', 'position'),
        ('486303cb-1b02-4059-a860-25704dcab8ea', 2, '/doctors', 'doctors_view', 'full_name', 'Специалисты', 'full_name', 'slug', 'doctor'),
        ('2665b300-222c-44c3-b33c-ad8babb71d93', 1, '/divisions', 'divisions', 'name', 'Отделения и центры', 'name', 'slug', 'division'),
        ('f68d3d90-60a0-45d1-a8fd-0d8d88a18eca', 3, '/paid-services', 'paid_services', 'name', null, 'name', 'id', 'paidService');