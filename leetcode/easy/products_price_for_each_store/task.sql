# Write your MySQL query statement below
select distinct a.product_id, b.store1, c.store2, d.store3
from products a
         left join
     (select product_id, price as store1
      from products
      where store = 'store1') b
     on a.product_id = b.product_id
         left join (select product_id, price as store2
                    from products
                    where store = 'store2') c
                   on a.product_id = c.product_id
         left join (select product_id, price as store3
                    from products
                    where store = 'store3') d
                   on a.product_id = d.product_id



