select
    stock_name,
    SUM(IF(operation = 'Buy', -price, price)) as capital_gain_loss
from Stocks
group by stock_name
