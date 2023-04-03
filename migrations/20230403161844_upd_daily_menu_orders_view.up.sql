drop view daily_menu_orders_view;

create
    or replace view daily_menu_orders_view as
SELECT
    dmo.*,
    fv.created_at,
    fv.is_new,
    fv.form_status_id,
    u.email,
    CONCAT_WS(' ' :: TEXT, h.surname, h.name, h.patronymic) AS full_name,
    fv.user_id
FROM
    daily_menu_orders dmo
        join form_values fv on fv.id = dmo.form_value_id
        join users u on u.id = fv.user_id
        join humans h on h.id = u.human_id;