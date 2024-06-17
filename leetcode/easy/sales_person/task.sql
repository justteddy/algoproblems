select name
from SalesPerson
where sales_id not in(
    select distinct t1.sales_id
    from Orders t1 left join Company t2 using(com_id)
    where t2.name = 'RED'
);
