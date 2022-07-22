CREATE TABLE news_views
(
    id         uuid DEFAULT uuid_generate_v4()                               NOT NULL PRIMARY KEY,
    news_id    uuid REFERENCES news (id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    ip_address VARCHAR                                                       NOT NULL,
    UNIQUE (news_id, ip_address)
);
