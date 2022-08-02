alter table users
    add created_at timestamp default current_timestamp not null;
