select t1.name, t2.balance
from users as t1 right join(
    select sum(amount) as balance, account
    from Transactions
    group by account
    having balance > 10000
) as t2 on t1.account = t2.account
