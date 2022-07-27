CREATE TABLE news_likes (
                              id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
                              news_id uuid  REFERENCES news (id) ON UPDATE CASCADE ON DELETE CASCADE,
                              user_id uuid  REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);
