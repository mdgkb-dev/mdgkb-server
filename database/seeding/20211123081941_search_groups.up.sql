insert into public.search_groups (id, key, search_group_order, route, search_group_table, search_column, label, label_column, value_column, description_column)
values  ('486303cb-1b02-4059-a860-25704dcab8ea', 'doctor', 2, '/doctors', 'doctors_view', 'full_name', 'Специалисты', 'full_name', 'slug', null),
        ('4f4cf39a-1125-4989-8c11-b4335517f31d', 'position', 4, '/positions', 'positions', 'name', 'Должности', 'name', 'id', null),
        ('f68d3d90-60a0-45d1-a8fd-0d8d88a18eca', 'paidService', 3, '/paid-services', 'paid_services', 'name', 'Платные услуги', 'name', 'id', null),
        ('76250b93-b54a-42a5-a6f8-3ea517960525', 'teacher', null, null, 'teachers_view', 'full_name', null, 'full_name', 'id', null),
        ('ffaf957c-ea44-4f47-8ae7-a817c58a0348', 'dpoCourse', null, null, 'dpo_courses', 'name', null, 'name', 'id', null),
        ('d8631b3a-79e9-4ec8-be98-58e77a199371', 'vacancy', null, '/vacancies', 'vacancies', 'title', null, 'title', 'slug', null),
        ('2665b300-222c-44c3-b33c-ad8babb71d93', 'division', 1, '/divisions', 'divisions_view', 'name', 'Отделения и центры', 'name', 'slug', 'info');