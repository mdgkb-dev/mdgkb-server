ALTER TABLE divisions
    ADD COLUMN treat_direction_id uuid  REFERENCES treat_directions(id) ON UPDATE CASCADE ON DELETE CASCADE;
