INSERT INTO public.humen (id, name, surname, patronymic, is_male, date_birth)
VALUES  ('fb40366b-b230-4e25-b3f7-46280c8eb983', 'Светлана', 'Короткова ', 'Анатольевна', false, '1995-09-07'),
        ('27746755-07f7-44ec-8ec3-e1a80a7cdba8', 'Екатерина', 'Моксякова', 'Геннадьевна', false, '1993-09-06'),
        ('6ec356e0-cd80-4617-b6dc-a98db98106d6', 'Елена', 'Полякова', 'Ивановна', false, '1994-09-22'),
        ('a9045308-ea24-4f41-a3b3-f3052c10426d', 'Гульноза', 'Тургунова', 'Мураджоновна', false, '1993-08-31'),
        ('be1622c0-ff7a-44c8-a529-04bb67687c5e', 'Гузель', 'Осипова', 'Тагировна', false, '1995-09-05'),
        ('e590dcee-29a7-448e-bb53-8d150e930dc6', 'Валерий', 'Горев', 'Викторович', true, '1978-08-03'),
        ('66058f82-41dc-4b86-bf37-373f0ce9d6cd', 'Ирина', 'Витковская', 'Петровна', false, '1991-08-31'),
        ('4c524c90-1181-4424-bb09-e5cd05c6025f', 'Мария', 'Шарапова', 'Дмитриевна', false, '1991-08-31');

INSERT INTO public.file_infos (id, original_name, file_system_path)
VALUES  ('a379fba7-7961-4932-baa6-f13913fa05a6', 'Короткова-Светлана-Анатольевна.png', 'ee9bb62b-0bdb-4d7c-87fa-1db1d5c43de6'),
        ('76e1eb1a-45d9-4d47-adcf-8291bd97a97d', 'Моксякова-Екатерина-Геннадьевна.png', '22ebd0f5-a5f1-4633-8afd-6c10a21b90a6'),
        ('52cc8ea7-8a4e-4f58-9875-15687ae9f938', 'Тургунова-Гульноза-Мураджоновна.png', 'b87696fc-7076-4785-aa41-7534b5c4d858'),
        ('a49c6c2e-7216-4608-8458-948b132964ee', 'Полякова-Елена-Ивановна.png', '5efafb3e-7011-4616-a698-d6a3d2d26553'),
        ('9c25c470-de38-4505-bf24-59cac0f7e7c9', 'Gorev_web.jpg', 'daa2778d-2e57-4df8-bc51-c76f0368167f'),
        ('49c56746-f7a6-4224-8295-fffec3cb6b82', 'Витковская-Ирина-Петровна.jpg', '52686728-faf8-47de-9a0e-cfc7ff96ed56'),
        ('051ffd03-5881-4ee9-984f-64f01ccd1f92', 'Осипова-Гузель-Тагировна.png', '6a2430fb-c50e-4c44-878f-e19fa86c1d74');

INSERT INTO public.doctors (id, human_id, division_id, education, schedule, position, tags, file_info_id)
VALUES  ('61802bb6-0b6e-4310-87f0-85d7652a487d', 'fb40366b-b230-4e25-b3f7-46280c8eb983', 'ef2093d9-0a8b-46e8-bf71-e42dad072090', 'Саратовский Ордена Трудового Красного Знамени государственный медицинский институт.', '9.00-18.00', 'Заведующий гинекологическим отделением - врачакушер-гинеколог', 'Гинекология', 'a379fba7-7961-4932-baa6-f13913fa05a6'),
        ('29730be2-f22e-4fbb-b015-14c65fb7184b', '27746755-07f7-44ec-8ec3-e1a80a7cdba8', 'ef2093d9-0a8b-46e8-bf71-e42dad072090', 'РГМУ', '9.00-18.00', 'Врач-акушергинеколог', 'Гинекология', '76e1eb1a-45d9-4d47-adcf-8291bd97a97d'),
        ('f7144dda-c1ae-4d5d-98ce-6a7abfae1e8c', '6ec356e0-cd80-4617-b6dc-a98db98106d6', 'ef2093d9-0a8b-46e8-bf71-e42dad072090', 'ФГБОУ ВПО "Орловский государственный университет"', '9.00-18.00', 'Врач-акушергинеколог', 'Гинекология', 'a49c6c2e-7216-4608-8458-948b132964ee'),
        ('8c280f14-0875-4388-b2cc-d77381bdbede', 'a9045308-ea24-4f41-a3b3-f3052c10426d', 'ef2093d9-0a8b-46e8-bf71-e42dad072090', 'Ошский государственный университет Кыргызская республика', '9.00-18.00', 'Врач-акушергинеколо', 'Гинекология', '52cc8ea7-8a4e-4f58-9875-15687ae9f938'),
        ('3f693d69-5822-49c5-bb0f-f4055ab47330', 'be1622c0-ff7a-44c8-a529-04bb67687c5e', 'ef2093d9-0a8b-46e8-bf71-e42dad072090', 'Башкирский ГМУ', '9.00-18.00', 'Врач-акушергинеколог', 'Гинекология', '051ffd03-5881-4ee9-984f-64f01ccd1f92'),
        ('ad36c2d6-403c-42c5-b4f9-7dbbd4c4bdba', 'e590dcee-29a7-448e-bb53-8d150e930dc6', NULL, 'Сибирский государственный медицинский университет г. Томска', '9.00-18.00', 'Главный врач, Главный внештатный специалист неонатолог, ', 'Главный врач', '9c25c470-de38-4505-bf24-59cac0f7e7c9'),
        ('a5cba2d5-6af1-4a02-a461-0ef6963907df', '66058f82-41dc-4b86-bf37-373f0ce9d6cd', NULL, 'Педиатрический факультет Российского государственного медицинского университета им. Н. И. Пирогова', '9.00-18.00', 'Заместитель главного врача по организационно-методической работе', 'Заместитель главного врача ', '49c56746-f7a6-4224-8295-fffec3cb6b82'),
        ('288771a1-2b98-418b-b8e0-276a75d9512f', '4c524c90-1181-4424-bb09-e5cd05c6025f', NULL, NULL, '9.00-18.00', 'Начальник отдела постдипломного образования', NULL, NULL);