CREATE TABLE pages_documents (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    document_id uuid REFERENCES documents (id),
    page_id uuid REFERENCES pages (id)
);