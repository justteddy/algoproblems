# Write your MySQL query statement below
select round(t1.immediate_percentage, 2) as immediate_percentage
from(
        select
            ((select count(delivery_id) from Delivery where order_date = customer_pref_delivery_date) /
             (select count(delivery_id) from Delivery)) * 100 as immediate_percentage
    ) t1
