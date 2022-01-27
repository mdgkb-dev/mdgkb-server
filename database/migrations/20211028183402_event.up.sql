create table events
(
    id         uuid default uuid_generate_v4() not null primary key,
    form_id    uuid references forms on update cascade on delete cascade,

    start_date timestamp,
    end_date   timestamp
);
