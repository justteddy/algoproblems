# Write your MySQL query statement below
select x.seller_id
from(
        select seller_id, sum(price) as total
        from Sales
        group by seller_id
        having total = (select sum(price) as total from Sales group by seller_id order by total desc limit 1)
order by total desc
    ) as x
