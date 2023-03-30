CREATE TABLE chats (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY
);

drop table chat_messages;
create table chat_messages
(
    id         uuid      default uuid_generate_v4() not null
        primary key,
    message    varchar,
    admin_id   uuid
        constraint chat_messages_user_id_fkey
            references users
            on update cascade on delete cascade,
    created_on timestamp default CURRENT_TIMESTAMP  not null,
    chat_id    uuid
        references chats
            on update cascade on delete cascade
);

alter table form_values
    add chat_id uuid;

alter table form_values
    add constraint form_values_chats_id_fk
        foreign key (chat_id) references chats
            on update cascade on delete cascade;

