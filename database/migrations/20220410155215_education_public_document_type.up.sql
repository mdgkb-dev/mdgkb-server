CREATE TABLE education_public_document_types
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    public_document_type_id  uuid  REFERENCES public_document_types(id) ON UPDATE CASCADE ON DELETE CASCADE
);
