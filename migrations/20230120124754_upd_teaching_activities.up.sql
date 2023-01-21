alter table teaching_activities
add column employee_id  uuid  REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;

update teaching_activities
set employee_id = e.id from employees e 
join doctors d on e.id = d.employee_id
where d.employee_id = e.id;

alter table teaching_activities drop  column doctor_id;
