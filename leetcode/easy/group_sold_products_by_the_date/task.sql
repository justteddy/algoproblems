# Write your MySQL query statement below
select t1.sell_date, count(t1.product) as num_sold, GROUP_CONCAT(t1.product order by t1.product) as products
from (
         select sell_date, product
         from Activities
         group by sell_date, product
     ) t1
GROUP BY sell_date
ORDER BY sell_date;
