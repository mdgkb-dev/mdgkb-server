CREATE TABLE pages_comments (
                                 id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
                                 comment_id uuid REFERENCES comments (id),
                                 page_id uuid REFERENCES pages (id)
);