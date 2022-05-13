CREATE TABLE educational_organization_document_types_documents (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    document_id uuid REFERENCES documents (id),
    educational_organization_document_type_id uuid REFERENCES educational_organization_document_types (id)
);