CREATE TABLE vacancy_responses (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    response_date date,
    cover_letter varchar,
    viewed boolean,
    vacancy_id uuid  REFERENCES vacancies (id) ON UPDATE CASCADE ON DELETE CASCADE,
    contact_info_id uuid references contact_infos(id) on update cascade on delete cascade,
    user_id uuid references users(id) on update cascade on delete cascade
);