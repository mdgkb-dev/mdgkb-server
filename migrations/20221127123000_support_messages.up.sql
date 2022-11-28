CREATE TABLE support_messages (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    question varchar,
    theme varchar,
    answer varchar,
    support_message_date date,
    is_new boolean,
    user_id uuid  REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);