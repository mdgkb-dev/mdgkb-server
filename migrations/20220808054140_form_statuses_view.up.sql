create or replace view form_statuses_view as
SELECT fsg.code, fs.*
FROM form_statuses fs
         JOIN form_status_groups fsg ON fs.form_status_group_id = fsg.id
group by fs.id, fsg.code;
