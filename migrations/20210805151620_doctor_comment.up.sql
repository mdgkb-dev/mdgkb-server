CREATE TABLE doctor_comments (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  doctor_id uuid,
  comment_id uuid 
);

