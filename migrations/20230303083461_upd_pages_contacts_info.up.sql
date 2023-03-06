alter table pages
    add show_contacts boolean default false;

alter table pages
    add contact_info_id uuid references contact_infos(id) on update cascade on delete cascade;
