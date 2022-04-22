CREATE OR REPLACE view applications_counts as
select count(fv.id) as new_applications, 'dpo_applications' as key from form_values as fv
        join dpo_applications da on fv.id = da.form_value_id
where fv.is_new = true
union
select count(fv.id) as new_applications, 'postgraduate_applications' as key from form_values as fv
        join postgraduate_applications pa on fv.id = pa.form_value_id
where fv.is_new = true
union
select count(fv.id) as new_applications, 'residency_applications' as key from form_values as fv
        join residency_applications ra on fv.id = ra.form_value_id
where fv.is_new = true