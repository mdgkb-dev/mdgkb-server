ALTER TABLE vacancy_responses
    ADD COLUMN form_value_id uuid  REFERENCES form_values (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;
