create table preparations_to_tags
(
    id uuid default uuid_generate_v4() not null primary key ,
    preparation_id uuid references preparations on delete cascade,
    preparation_tag_id uuid references preparations_tags on delete cascade
);

