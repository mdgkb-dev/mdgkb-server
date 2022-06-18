-- drop view applications_cars_view;

create
or replace view applications_cars_view as
SELECT
    ac.*,
    fv.created_at,
    fv.form_status_id,
    gates.name as gate_name,
    divisions.name as division_name,
    u.email,
    CONCAT_WS(' ' :: TEXT, h.surname, h.name, h.patronymic) AS child_full_name
FROM
    applications_cars ac
    join form_values fv on fv.id = ac.form_value_id
    join gates on gates.id = ac.gate_id
    join divisions on divisions.id = ac.division_id
    join users u on u.id = fv.user_id
    join children on children.id = fv.child_id
    join humans h on h.id = children.human_id;