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

--drop table normative_documents;
--drop table normative_document_types;

