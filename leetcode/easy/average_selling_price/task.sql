# Write your MySQL query statement below
select t.product_id, IFNULL(round(sum(t.unit_price * t.units) / sum(t.units), 2), 0) as average_price
from(
        select t2.product_id, t1.purchase_date, t1.units, t2.price as unit_price
        from Prices t2
                 left join UnitsSold t1 on
            t1.product_id = t2.product_id and t1.purchase_date between t2.start_date and t2.end_date
    ) t
group by t.product_id;
