# Write your MySQL query statement below
select t2.product_name, t1.year, t1.price
from Sales t1 left join Product t2 Using(product_id);
