CREATE TABLE dpo_applications (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    created_at timestamp,
    dpo_course_id uuid REFERENCES dpo_courses (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    user_id uuid REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);