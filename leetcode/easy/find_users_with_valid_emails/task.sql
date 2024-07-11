# Write your MySQL query statement below
select user_id, name, mail
from Users
where mail REGEXP '^[A-Za-z]{1}[A-Za-z0-9_\.-]{0,}@leetcode[\.]{1}com$'