CREATE TABLE news_comments (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  news_id uuid REFERENCES news (id),
  comment_id uuid 
    constraint news_comments_comments_id_fk 
    REFERENCES comments 
    on delete cascade
);

create unique index news_comments_id_uindex
    on news_comments (id);
