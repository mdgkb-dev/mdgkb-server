CREATE TABLE medical_profiles_news
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    medical_profile_id uuid  REFERENCES medical_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    news_id uuid  REFERENCES news (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);