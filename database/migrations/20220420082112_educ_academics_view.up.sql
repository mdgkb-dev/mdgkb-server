create
or replace view educational_organization_academics_view as
SELECT
    dv.full_name,
    ea.*
FROM
    educational_organization_academics ea
    JOIN doctors_view dv ON dv.id = ea.doctor_id
group by
    ea.id, dv.full_name;