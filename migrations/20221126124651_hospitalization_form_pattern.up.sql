ALTER TABLE hospitalizations_types
    ADD COLUMN form_pattern_id uuid REFERENCES form_patterns(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE hospitalizations_types
    ADD COLUMN hospitalization_type_order integer default 0;

