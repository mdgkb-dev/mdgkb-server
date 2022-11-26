ALTER TABLE hospitalizations_types
    ADD COLUMN form_pattern_id uuid REFERENCES form_patterns(id) ON UPDATE CASCADE ON DELETE CASCADE;
