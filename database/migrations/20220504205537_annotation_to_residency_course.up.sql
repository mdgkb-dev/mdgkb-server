ALTER TABLE postgraduate_courses
    ADD COLUMN annotation_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE;
