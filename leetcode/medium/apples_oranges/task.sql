# Write your MySQL query statement below
select sale_date, sum(sold_num) as diff
from(
        select sale_date, fruit, if(fruit='apples', sold_num, -sold_num) as sold_num
        from sales
    ) as t
group by sale_date;
