# Write your MySQL query statement below
select t1.id, t1.name
from Students as t1 left join Departments as t2 on t1.department_id = t2.id
where t2.id is null;
