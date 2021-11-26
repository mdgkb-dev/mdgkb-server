CREATE TABLE faqs
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    question varchar,
    answer varchar
);
