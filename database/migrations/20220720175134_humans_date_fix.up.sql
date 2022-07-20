alter table humans
ALTER COLUMN date_birth TYPE timestamp;

update humans
    set date_birth = date_birth + interval '1h' * 21
where date_birth is not null