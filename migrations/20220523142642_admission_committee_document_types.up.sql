CREATE TABLE admission_committee_document_types
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    document_type_id  uuid  REFERENCES document_types(id) ON UPDATE CASCADE ON DELETE CASCADE
);
