ALTER TABLE visits_applications ADD COLUMN with_car boolean default false;

drop view visits_applications_view;

create
or replace view visits_applications_view as
SELECT
    va.*,
    fv.created_at,
    fv.form_status_id,
    gates.name as gate_name,
    divisions.name as division_name,
    u.email,
    CONCAT_WS(' ' :: TEXT, h.surname, h.name, h.patronymic) AS child_full_name
FROM
    visits_applications va
    join form_values fv on fv.id = va.form_value_id
    join gates on gates.id = va.gate_id
    join divisions on divisions.id = va.division_id
    join users u on u.id = fv.user_id
    join children on children.id = fv.child_id
    join humans h on h.id = children.human_id;