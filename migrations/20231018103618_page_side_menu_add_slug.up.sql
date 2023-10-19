alter table page_side_menus add column slug varchar;


select * from page_side_menus;

update page_side_menus set slug = 'stati-o-bolnicze' where name='Статьи о больнице';
update page_side_menus set slug = 'priyomnaya-kampaniya' where name='Приёмная кампания';
update page_side_menus set slug = 'programmy' where name='Программы';
update page_side_menus set slug = 'chasto-zadavaemye-voprosy' where name='Часто задаваемые вопросы';
update page_side_menus set slug = 'obrazovanie' where name='Образование';
update page_side_menus set slug = 'obrazovatelnye-standarty' where name='Образовательные стандарты';
update page_side_menus set slug = 'materialno-tekhnicheskoe-obespechenie-i-osnashhennost-obrazovatelnogo-proczessa' where name='Материально-техническое обеспечение и оснащенность образовательного процесса';
update page_side_menus set slug = 'platnye-obrazovatelnye-uslugi' where name='Платные образовательные услуги';
update page_side_menus set slug = 'soglasie-na-zachislenie-i-zaklyuchenie-dogovorov' where name='Согласие на зачисление и заключение договоров';
update page_side_menus set slug = 'dokumenty' where name='Документы';
update page_side_menus set slug = 'osnovnye-svedeniya' where name='Основные сведения';
update page_side_menus set slug = 'mezhdunarodnoe-sotrudnichestvo' where name='Международное сотрудничество';
update page_side_menus set slug = 'bolnicza-segodnya' where name='Больница сегодня';
update page_side_menus set slug = 'istoriya' where name='История';
update page_side_menus set slug = 'video-o-bolnicze' where name='Видео о больнице';
update page_side_menus set slug = 'normativnye-dokumenty' where name='Нормативные документы';
update page_side_menus set slug = 'data-i-vremya-priema' where name='Дата и время приема';
update page_side_menus set slug = 'normativno-pravovaya-baza' where name='Нормативно-правовая база';
update page_side_menus set slug = 'informacziya-dlya-roditelej' where name='Информация для родителей';
update page_side_menus set slug = 'informacziya-dlya-mediczinskogo-personala' where name='Информация для медицинского персонала';
update page_side_menus set slug = 'informacziya-o-czentre' where name='Информация о Центре';
update page_side_menus set slug = 'dokumenty-priyomnoj-komissii' where name='Документы приёмной комиссии';
update page_side_menus set slug = 'informacziya-o-priyomnoj-kampanii' where name='Информация о приёмной кампании';
update page_side_menus set slug = 'plan-nabora' where name='План набора';
update page_side_menus set slug = 'stazhirovka' where name='Стажировка';
update page_side_menus set slug = 'povyshenie-kvalifikaczii' where name='Повышение квалификации';
update page_side_menus set slug = 'dostupnaya-sreda' where name='Доступная среда';
update page_side_menus set slug = 'okazyvaemye-uslugi' where name='Оказываемые услуги';
update page_side_menus set slug = 'informacziya-dlya-inostrannykh-grazhdan' where name='Информация для иностранных граждан';
update page_side_menus set slug = 'informacziya-ob-uchyote-individualnykh-dostizhenij-bally-nachislyayutsya-odnokratno-za-kazhdoe-dostizhenie' where name='Информация об учёте индивидуальных достижений (баллы начисляются однократно за каждое достижение)';
update page_side_menus set slug = 'obemy-obyazatelnykh-laboratornykh-i-funkczionalnykh-issledovanij-pri-provedenii-obyazatelnogo-predvaritelnogo-mediczinskogo-osmotra-pri-postuplenii-v-ordinaturu-' where name='Объемы обязательных лабораторных и функциональных исследований при проведении обязательного предварительного медицинского осмотра при поступлении в ординатуру ';
update page_side_menus set slug = 'informacziya-ob-obshhezhitii' where name='Информация об общежитии';
update page_side_menus set slug = 'speczialisty' where name='Специалисты';
update page_side_menus set slug = 'kontakty' where name='Контакты';
update page_side_menus set slug = 'osnovye-svedeniya' where name='Основые сведения';
update page_side_menus set slug = 'finansovo-khozyajstvennaya-deyatelnost' where name='Финансово-хозяйственная деятельность';
update page_side_menus set slug = 'informacziya-dlya-paczientov-s-sma' where name='Информация для пациентов с СМА';
update page_side_menus set slug = 'raznoe' where name='Разное';
update page_side_menus set slug = 'dokumenty-dlya-postupayushhikh-v-ordinaturu' where name='Документы для поступающих в ординатуру';
update page_side_menus set slug = 'stipendii-i-inye-vidy-materialnoj-podderzhki' where name='Стипендии и иные виды материальной поддержки';
update page_side_menus set slug = 'vakantnye-mesta-dlya-priyoma-perevoda' where name='Вакантные места для приёма (перевода)';
update page_side_menus set slug = 'vstupitelnye-ispytaniya' where name='Вступительные испытания';
update page_side_menus set slug = 'pamyatka-dlya-posetitelej-paczientov-staczionara' where name='Памятка для посетителей пациентов стационара';
update page_side_menus set slug = 'rukovodstvo-pedagogicheskij-sostav' where name='Руководство. Педагогический состав';
update page_side_menus set slug = 'obshhaya-informacziya' where name='Общая информация';
update page_side_menus set slug = 'vydacha-tekhnicheskikh-sredstv-reabilitaczii-tsr' where name='Выдача технических средств реабилитации (ТСР)';
update page_side_menus set slug = 'dispanserizacziya-naseleniya' where name='Диспансеризация населения';
update page_side_menus set slug = 'pamyatki-pri-lyogkom-techenii-covid-19-i-orvi' where name='Памятки при лёгком течении COVID-19 и ОРВИ';
update page_side_menus set slug = 'uchreditelnye-dokumenty' where name='Учредительные документы';
update page_side_menus set slug = 'protivodejstvie-korrupczii' where name='Противодействие коррупции';
update page_side_menus set slug = 'czelevaya-ordinatura' where name='Целевая ординатура';
update page_side_menus set slug = 'ordinatura-po-dogovoram-o-platnykh-uslugakh' where name='Ординатура по договорам о платных услугах';
update page_side_menus set slug = 'prikazy-na-zachislenie-informacziya-o-01092023' where name='Приказы на зачисление. Информация о 01.09.2023';
update page_side_menus set slug = 'czelevaya-ordinatura' where name='Целевая ординатура';
update page_side_menus set slug = 'ordinatura-po-dogovoram-o-platnykh-obrazovatelnykh-uslugakh' where name='Ординатура по договорам о платных образовательных услугах';
update page_side_menus set slug = 'svedeniya-o-mediczinskoj-organizaczii' where name='Сведения о медицинской организации';

