CREATE TABLE news_divisions
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    news_id uuid  REFERENCES news (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
        UNIQUE (news_id, division_id)
);