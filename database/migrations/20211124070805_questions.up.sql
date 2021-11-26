CREATE TABLE questions (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    question varchar,
    answer varchar,
    original_question varchar,
    original_answer varchar,
    published boolean,
    question_date date,
    is_new boolean,
    answer_is_read boolean,
    user_id uuid  REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);