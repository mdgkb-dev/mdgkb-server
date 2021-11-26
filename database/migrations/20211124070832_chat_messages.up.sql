CREATE TABLE chat_messages (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    message varchar,
    is_answer boolean,
    chat_message_date date,
    user_id uuid  REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    answer_user_id uuid  REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);