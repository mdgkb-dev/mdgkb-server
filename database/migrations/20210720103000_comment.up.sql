CREATE TABLE comments (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  user_id uuid constraint comments_user_id_fk REFERENCES users,
  text TEXT,
  rating float,
  mod_checked boolean,
  positive boolean,
  published_on timestamp default current_timestamp
);

create unique index comments_id_uindex on comments (id);