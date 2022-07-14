CREATE
OR REPLACE view applications_counts as
select
    count(fv.id) as count,
    'dpo_applications' as table_name
from
    form_values as fv
    join dpo_applications da on fv.id = da.form_value_id
    join dpo_courses dc on da.dpo_course_id = dc.id
where
    dc.is_nmo = false
    and fv.is_new = true
union
select
    count(fv.id) as count,
    'nmo_applications' as table_name
from
    form_values as fv
    join dpo_applications da on fv.id = da.form_value_id
    join dpo_courses dc on da.dpo_course_id = dc.id
where
    dc.is_nmo = true
    and fv.is_new = true
union
select
    count(fv.id) as count,
    'postgraduate_applications' as table_name
from
    form_values as fv
    join postgraduate_applications pa on fv.id = pa.form_value_id
where
    fv.is_new = true
union
select
    count(fv.id) as count,
    'residency_applications' as table_name
from
    form_values as fv
    join residency_applications ra on fv.id = ra.form_value_id
where
    fv.is_new = true
union
select
    count(comments.id) as count,
    'comments' as table_name
from
    comments
where
    comments.mod_checked = false
union
select
    count(fv.id) as count,
    'vacancy_responses' as table_name
from
    form_values as fv
    join vacancy_responses vr on fv.id = vr.form_value_id
where
    fv.is_new = true
union
select
    count(fv.id) as count,
    'visits_applications' as table_name
from
    form_values as fv
    join visits_applications ac on fv.id = ac.form_value_id
where
    fv.is_new = true