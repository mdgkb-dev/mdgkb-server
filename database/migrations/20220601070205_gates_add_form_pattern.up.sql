ALTER TABLE gates
    ADD COLUMN form_pattern_id uuid references form_patterns on update cascade on delete cascade;
