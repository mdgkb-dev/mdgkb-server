CREATE TABLE news_to_categories (
     id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
     news_id uuid  REFERENCES news (id) ON UPDATE CASCADE ON DELETE CASCADE,
     category_id uuid REFERENCES categories (id) ON UPDATE CASCADE ON DELETE CASCADE
);
