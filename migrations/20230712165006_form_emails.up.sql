CREATE TABLE form_status_emails (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    email VARCHAR,
    form_status_id uuid REFERENCES form_statuses(id) ON UPDATE CASCADE ON DELETE CASCADE,
    template_file_name varchar,
    theme varchar
);


insert into form_status_emails (id, email, form_status_id, template_file_name, theme)
values  ('15d41912-9daa-4905-932e-85b162dfe60c', 'lakkinzimusic@gmail.com', '03977c2f-31a9-4160-b194-5e595d89e434', 'residencyApplicationAdmin.gohtml', 'Новая заявка в ординатуру');