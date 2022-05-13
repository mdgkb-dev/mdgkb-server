CREATE TABLE event_applications (
                        id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
                        event_id uuid  REFERENCES events (id) ON UPDATE CASCADE ON DELETE CASCADE,
                        user_id uuid  REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

