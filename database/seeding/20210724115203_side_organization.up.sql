do $$
declare
   contact_info_id_01 uuid = uuid_generate_v4();
   contact_info_id_02 uuid = uuid_generate_v4();
   contact_info_id_03 uuid = uuid_generate_v4();
   contact_info_id_04 uuid = uuid_generate_v4();
   contact_info_id_05 uuid = uuid_generate_v4();
   contact_info_id_06 uuid = uuid_generate_v4();
   contact_info_id_07 uuid = uuid_generate_v4();
   contact_info_id_08 uuid = uuid_generate_v4();
   contact_info_id_09 uuid = uuid_generate_v4();
   contact_info_id_10 uuid = uuid_generate_v4();
   contact_info_id_11 uuid = uuid_generate_v4();
   contact_info_id_12 uuid = uuid_generate_v4();
   contact_info_id_13 uuid = uuid_generate_v4();
   contact_info_id_14 uuid = uuid_generate_v4();
   contact_info_id_15 uuid = uuid_generate_v4();
   contact_info_id_16 uuid = uuid_generate_v4();
   contact_info_id_17 uuid = uuid_generate_v4();
   contact_info_id_18 uuid = uuid_generate_v4();
begin
   insert into
      public.contact_infos (id)
   values
      (contact_info_id_01),
      (contact_info_id_02),
      (contact_info_id_03),
      (contact_info_id_04),
      (contact_info_id_05),
      (contact_info_id_06),
      (contact_info_id_07),
      (contact_info_id_08),
      (contact_info_id_09),
      (contact_info_id_10),
      (contact_info_id_11),
      (contact_info_id_12),
      (contact_info_id_13),
      (contact_info_id_14),
      (contact_info_id_15),
      (contact_info_id_16),
      (contact_info_id_17),
      (contact_info_id_18);

   insert into
      public.side_organizations (name, description, contact_info_id)
   values
      ('Единая справочная службы г. Москвы (в том числе по вопросам доступности и качества бесплатной медицинской помощи)', null, contact_info_id_01),
      ('Госпитализация иногородних пациентов', null, contact_info_id_02),
      ('Справочная служба по вопросам лекарственного обеспечения', 'Часы работы: пн.-сб. с 8:00 до 20:00, вс. - выходной день', contact_info_id_03),
      ('Справочная служба по вопросам применения цен на лекарственные препараты, включенные в перечень жизненно необходимых и важнейших лекарственных препаратов', 'Часы работы: пн.-чт. с 9:00 до 17:45, пт. с 9:00 до 16:30, кроме праздничных дней. Обеденный перерыв: 13:30-14:00', contact_info_id_04),
      ('Департамент здравоохранения города Москвы', 'Часы работы: пн.–пт. c 08.00 до 20.00, сб. и вс. – выходные дни', contact_info_id_05),
      ('Оперативно-распорядительная служба Департамента здравоохранения города Москвы', null, contact_info_id_06),
      ('Дежурный врач Станции скорой и неотложной медицинской помощи им. А.С. Пучкова', null, contact_info_id_07),
      ('Справка о госпитализации больных', 'В зависимости от места проживания можно также обращаться по справочным телефонам Дирекции по координации деятельности медицинских организаций Департамента здравоохранения города Москвы и Дирекции по обеспечению деятельности государственных учреждений здравоохранения Троицкого и Новомосковского административных округов', contact_info_id_08),
      ('Дирекция по координации деятельности медицинских организаций', null, contact_info_id_09),
      ('Дирекция по обеспечению деятельности государственных учреждений здравоохранения Троицкого и Новомосковского административных округов', null, contact_info_id_10),
      ('Телефон психологической помощи в Москве', null, contact_info_id_11),
      ('Московский городской фонд обязательного медицинского страхования', null, contact_info_id_12),
      ('Департамент труда и социальной защиты населения города Москвы', null, contact_info_id_13),
      ('Территориальный орган Федеральной службы по надзору в сфере здравоохранения (Росздравнадзор) по г. Москве и Московской области', 'Время работы: пн.–чт. с 9:00 до 17:45, пт. с 9:00 до 16:30, сб. и вс. – выходные дни', contact_info_id_14),
      ('Управление Федеральной службы по надзору в сфере защиты прав потребителей и благополучия человека по городу Москве (Управление Роспотребнадзора по г. Москве)', 'Почтовый адрес: 129626, г. Москва, Графский переулок, д.4/9. Время приёма обращения граждан: пн.-чт. c 9:00 до 17:30, пт. c 9:00 до 16:30. Обед с 13:00 до 13:45', contact_info_id_15),
      ('Министерство здравоохранения Росcийской Федерации', 'Почтовый адрес: 127994, ГСП-4, г. Москва, Рахмановский пер, д.3', contact_info_id_16),
      ('Федеральное казенное учреждение «Главное бюро медико-социальной экспертизы по г. Москве» Министерства труда и социальной защиты Российской Федерации (ФКУ «ГБ МСЭ по г. Москве» Минтруда России)', null, contact_info_id_17),
      ('Дирекция по координации деятельности медицинских организаций Департамента здравоохранения города Москвы', 'Режим работы: пн.–чт. с 9.00 до 17.45, пт. с 9.00 до 16.45, сб. и вс. – выходные дни. Почтовый адрес: Москва, 2-й Автозаводский проезд, дом 3', contact_info_id_18);

      insert into
         public.post_addresses (address, description, contact_info_id)
      values
         ('127006, г. Москва, Оружейный переулок, д.43', null, contact_info_id_05),
         ('127473, г. Москва, ул. Достоевского, д. 31, корп. 1А', null, contact_info_id_12),
         ('127206, г. Москва, ул. Вучетича, д.12а', null, contact_info_id_14),
         ('129626, г. Москва, Графский переулок, д.4/9', null, contact_info_id_15),
         ('127994, ГСП-4, г. Москва, Рахмановский пер, д.3', null, contact_info_id_16),
         ('125040, г. Москва, Ленинградский пр-т, д. 13, стр. 1', null, contact_info_id_17),
         ('115280, г. Москва, 2-й Автозаводский проезд, дом 3', null, contact_info_id_18);

      insert into
         public.telephone_numbers (number, description, contact_info_id)
      values
         ('+7 (495) 777-77-77', null, contact_info_id_01),
         ('+7 (495) 587-70-88', null, contact_info_id_02),
         ('+7 (495) 974-63-65', null, contact_info_id_03),
         ('+7 (495) 974-63-65', null, contact_info_id_03),
         ('+7 (495) 531-69-89', null, contact_info_id_04),
         ('+7 (499) 194-27-74', 'Горячая линия по вопросам вакцинации', contact_info_id_05),
         ('+7 (499) 251-44-27', 'Дежурный пост (круглосуточно)', contact_info_id_05),
         ('+7 (499) 251-83-00', '(круглосуточно)', contact_info_id_06),
         ('103', 'Общий телефон', contact_info_id_07),
         ('+7 (495) 620-42-33', 'Врачебно-консультативный пульт (дежурный врач)', contact_info_id_07),
         ('+7 (495) 620-42-25', 'Дежурный врач-педиатр', contact_info_id_07),
         ('+7 (495) 620-41-40', null, contact_info_id_08),
         ('+7 (495) 318-00-11', null, contact_info_id_09),
         ('+7 (495) 318-00-11', 'ВАО', contact_info_id_09),
         ('+7 (495) 439-44-02', 'ЗАО', contact_info_id_09),
         ('+7 (495) 946-11-00', 'САО', contact_info_id_09),
         ('+7 (495) 946-11-09', 'САО', contact_info_id_09),
         ('+7 (495) 610-65-20', 'СВАО', contact_info_id_09),
         ('+7 (499) 198-55-10', 'СЗАО', contact_info_id_09),
         ('+7 (495) 318-47-71', 'ЮАО', contact_info_id_09),
         ('+7 (495) 530-12-76', 'ЮВАО', contact_info_id_09),
         ('+7 (499) 125-62-00', 'ЮЗАО', contact_info_id_09),
         ('+7 (495) 951-67-65', 'ЦАО', contact_info_id_09),
         ('+7 (499) 734-11-91', 'ЗелАО', contact_info_id_09),
         ('+7 (499) 731-90-03', 'ЗелАО', contact_info_id_09),
         ('+7 (499) 391-35-90', 'ТИНАО', contact_info_id_09),
         ('+7 (499) 347-06-16', '(круглосуточно)', contact_info_id_10),
         ('051', 'С городского телефона - бесплатно. С мобильного телефона оплачиваются только услуги оператора связи согласно тарифному плану', contact_info_id_11),
         ('+7 (495) 623-10-20', '(по понедельникам с 15:00 до 18:00)', contact_info_id_12),
         ('+7 (495) 952-93-21', null, contact_info_id_12),
         ('+7 (495) 958-18-08', '(факс)', contact_info_id_12),
         ('+7 (495) 623-10-20', '(по понедельникам с 15:00 до 18:00)', contact_info_id_13),
         ('+7 (495) 611-47-74', null, contact_info_id_14),
         ('+7 (916) 256-76-76', null, contact_info_id_14),
         ('+7 (495) 318-47-71', null, contact_info_id_15),
         ('+7 (495) 687-40-35', null, contact_info_id_15),
         ('+7 (495) 616-65-69', '(факс)', contact_info_id_15),
         ('+7 (495) 628-44-53', 'Справочная служба', contact_info_id_16),
         ('+7 (495) 627-29-44', 'Справочная служба', contact_info_id_16),
         ('+7 (495) 627-24-00', 'Многоканальный телефон', contact_info_id_16),
         ('+7 (495) 916-03-09', 'Горячая линия', contact_info_id_17),
         ('+7 (495) 318-00-11', null, contact_info_id_18);

      insert into
         public.websites (address, description, contact_info_id)
      values
         ('mosgorzdrav.ru', null, contact_info_id_05),
         ('mgfoms.ru', null, contact_info_id_12),
         ('77reg.roszdravnadzor.ru', null, contact_info_id_14),
         ('77reg.roszdravnadzor.ru', null, contact_info_id_15),
         ('rosminzdrav.ru', null, contact_info_id_16),
         ('gbmsem.ru', null, contact_info_id_17),
         ('dkdmozdrav.ru', null, contact_info_id_18);
end $$;
