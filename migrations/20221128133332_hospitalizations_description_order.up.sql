ALTER TABLE hospitalizations_types
    ADD COLUMN description varchar;

ALTER TABLE hospitalization_type_analyzes
    ADD COLUMN item_order integer;

ALTER TABLE hospitalization_type_stages
    ADD COLUMN item_order integer;

ALTER TABLE hospitalization_type_documents
    ADD COLUMN item_order integer;