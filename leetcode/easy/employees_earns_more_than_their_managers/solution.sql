-- https://leetcode.com/problems/employees-earning-more-than-their-managers/
select E.Name as Employee
from Employee as E
         LEFT JOIN Employee as M ON E.ManagerId = M.Id
WHERE E.Salary > M.Salary
