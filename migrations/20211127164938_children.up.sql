CREATE TABLE children (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    human_id uuid  REFERENCES humans (id) ON UPDATE CASCADE ON DELETE CASCADE,
    user_id uuid  REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);