CREATE TABLE doctor_comments (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  doctor_id uuid REFERENCES doctors (id),
  comment_id uuid 
    constraint doctor_comments_comments_id_fk 
    REFERENCES comments 
    on delete cascade
);

create unique index doctor_comments_id_uindex
    on doctor_comments (id);
