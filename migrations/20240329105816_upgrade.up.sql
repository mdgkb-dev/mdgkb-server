alter table telephone_numbers rename to phones;
alter table contact_infos rename to contacts;

alter table pages rename column contact_info_id to contact_id;
alter table emails rename column contact_info_id to contact_id;
alter table phones rename column contact_info_id to contact_id;
alter table post_addresses rename column contact_info_id to contact_id;
alter table websites rename column contact_info_id to contact_id;
alter table address_infos rename to address;

alter table humans rename column contact_info_id to contact_id;





