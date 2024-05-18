select t1.name as warehouse_name, sum(t1.units * t2.unitVolume) as volume
from Warehouse as t1 left join (
    select product_id, Width*Height*Length as unitVolume
    from Products
) as t2 on t1.product_id = t2.product_id
group by warehouse_name
