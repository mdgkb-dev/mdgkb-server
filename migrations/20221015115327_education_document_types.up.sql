alter table documents
    add download_to_file bool default false;

    alter table public_document_types
        add public_document_type_order int;



CREATE TABLE document_types_images (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  description varchar,
  document_type_image_order int,
  document_type_id uuid REFERENCES document_types (id) ON UPDATE CASCADE ON DELETE CASCADE,
  file_info_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE
);
