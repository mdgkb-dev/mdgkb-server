CREATE TABLE news_doctors
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    doctor_id uuid  REFERENCES doctors (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    news_id uuid  REFERENCES news (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
        UNIQUE (news_id, doctor_id)
);