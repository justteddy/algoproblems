# Write your MySQL query statement below
select bin, total
from(
        select '[0-5>' as bin, count(session_id) as total
        from Sessions
        where duration < 300
        UNION
        select '[5-10>' as bin, count(session_id) as total
        from Sessions
        where duration >= 300 and duration < 600
        UNION
        select '[10-15>' as bin, count(session_id) as total
        from Sessions
        where duration >= 600 and duration < 900
        UNION
        select '15 or more' as bin, count(session_id) as total
        from Sessions
        where duration >= 900
    ) t1
