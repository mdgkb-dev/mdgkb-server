create or replace view users_view as
SELECT
    users.*,
    h.name,
    h.surname,
    h.patronymic,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM users
        join humans h on h.id = users.human_id
;
