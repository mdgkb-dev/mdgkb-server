CREATE TABLE comments (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  user_id uuid constraint comments_user_id_fk REFERENCES users,
  text TEXT,
  published_on date
);

create unique index comments_id_uindex
    on comments (id);
