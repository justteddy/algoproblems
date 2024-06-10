# Write your MySQL query statement below
select t1.product_name, sum(t2.unit) as unit
from Products t1 left join Orders t2 USING(product_id)
where DATE_FORMAT(t2.order_date, '%Y-%m') = '2020-02'
group by t1.product_name
having unit >= 100;
