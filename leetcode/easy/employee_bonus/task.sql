# Write your MySQL query statement below
select t1.name, t2.bonus
from Employee t1 left join Bonus t2 using(empId)
where bonus is null or bonus < 1000;
