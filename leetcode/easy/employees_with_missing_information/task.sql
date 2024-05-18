select t1.employee_id
from Employees as t1 left join Salaries as t2 on t1.employee_id = t2.employee_id
where t2.salary is null
union
select t1.employee_id
from Salaries as t1 left join Employees as t2 on t1.employee_id = t2.employee_id
where t2.name is null
order by employee_id;
