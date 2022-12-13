alter table public_document_types
    rename to page_side_menus;

alter table page_side_menus
    rename column public_document_type_order to item_order;

alter table page_side_menus
    ADD COLUMN page_id uuid  REFERENCES pages(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;


alter table document_types
    rename to page_sections;

alter table page_sections
    ADD COLUMN page_id uuid  REFERENCES pages(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

alter table page_sections
    ADD COLUMN page_side_menu_id uuid  REFERENCES page_side_menus(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

alter table page_sections
    rename column document_type_order  to item_order;

alter table document_types_images
    rename to page_section_images;

drop table document_field_values;
drop table document_type_fields;
drop table documents_types_for_vacancies;

alter table page_sections
    ADD COLUMN page_side_menu_id uuid  REFERENCES page_side_menus(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

alter table documents
    rename to page_section_documents;

alter table page_section_documents
    ADD COLUMN scan_id uuid  REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

update page_section_documents set
 scan_id = ds.scan_id from documents_scans ds
where ds.document_id = page_section_documents.id;

drop table documents_scans;

drop table normative_documents;
drop table normative_document_types;

alter table page_section_documents
    rename column document_type_id to page_section_id;

alter table page_section_documents
    rename column document_order to item_order;




    update page_side_menus set page_id = (select id from pages where slug = 'educational-info')
    where page_side_menus.id in (
    select psm.id from education_public_document_types
             join page_side_menus psm on psm.id = education_public_document_types.public_document_type_id
             where public_document_type_id is not null);

    drop table education_public_document_types;

    insert into page_side_menus (id, name, description, item_order)
    select ps.id, ps.name, ps.description, acdt.admission_committee_document_type_order
    from page_sections ps
    join admission_committee_document_types acdt on ps.id = acdt.document_type_id;


    update page_side_menus set page_id = (select id from pages where slug = 'admission-committee')
    where page_side_menus.id in (
    select psm.id from admission_committee_document_types
             join page_side_menus psm on psm.id = admission_committee_document_types.document_type_id
             where admission_committee_document_types.document_type_id is not null);

    drop table admission_committee_document_types;

   update page_sections set page_id = (select id from pages where slug = 'postgraduate')
    where page_sections.id in (
    select document_type_id  from postgraduate_document_types
             where postgraduate_document_types.document_type_id is not null);

    drop table postgraduate_document_types;

insert into page_side_menus (name, page_id)
    select 'Кандидатский минимум', id from pages where slug = 'postgraduate' ;

    update page_sections set page_side_menu_id = (select id from page_side_menus where name = 'Кандидатский минимум')
    where page_sections.id in (
    select candidate_document_types.document_type_id from candidate_document_types
             where candidate_document_types.document_type_id is not null);

drop table candidate_document_types;
drop table dpo_document_types;
drop table residency_document_types;

    update page_sections
    set page_side_menu_id = public_document_type_id
    where id is not null;



