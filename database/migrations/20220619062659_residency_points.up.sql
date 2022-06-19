alter table residency_applications
add column points_achievements int,
add column points_entrance int;

drop view residency_applications_view;

create
    or replace view residency_applications_view as
SELECT
    ra.*,
    ra.points_achievements + ra.points_entrance as points_sum,
    ey_s.year as start_year,
    ey_e.year as end_year,
    fv.created_at,
    fv.form_status_id,
    u.email,
    CONCAT_WS(' ' :: TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM
    residency_applications ra
        join form_values fv on fv.id = ra.form_value_id
        join residency_courses rc on rc.id = ra.residency_course_id
        join education_years ey_s on ey_s.id = rc.start_year_id
        join education_years ey_e on ey_e.id = rc.end_year_id
        join users u on u.id = fv.user_id
        join humans h on h.id = u.human_id;
