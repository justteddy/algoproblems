# Write your MySQL query statement below
select customer_number
from(
        select customer_number, count(customer_number) as total
        from Orders
        group by customer_number
        order by total desc) t1
    limit 1
