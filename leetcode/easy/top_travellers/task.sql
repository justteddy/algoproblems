# Write your MySQL query statement below
select t1.name, IFNULL(sum(t2.distance), 0) as travelled_distance
from users t1 left join rides t2 on t1.id = t2.user_id
group by t1.id
order by travelled_distance desc, t1.name asc;
