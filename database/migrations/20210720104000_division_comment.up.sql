CREATE TABLE division_comments (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  division_id uuid REFERENCES divisions (id),
  comment_id uuid 
    constraint division_comments_comments_id_fk 
    REFERENCES comments 
    on delete cascade
);

create unique index division_comments_id_uindex
    on division_comments (id);
