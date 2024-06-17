select customer_id, name
from(
        select t.customer_id, count(distinct t.month) as unique_months, c.name
        from(
                select t1.customer_id, DATE_FORMAT(t1.order_date, '%Y-%m') AS month, sum(t2.price*t1.quantity) as total
                from orders t1 left join Product t2 using(product_id)
                where DATE_FORMAT(t1.order_date, '%Y-%m') in ('2020-06', '2020-07')
                group by t1.customer_id, DATE_FORMAT(t1.order_date, '%Y-%m')
                having total >= 100
            ) as t left join Customers c using(customer_id)
        group by customer_id
        having unique_months = 2
    ) x
