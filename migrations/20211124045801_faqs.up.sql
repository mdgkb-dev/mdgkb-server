CREATE TABLE faqs
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    faq_order int not null default 0,
    question varchar,
    answer varchar
);
