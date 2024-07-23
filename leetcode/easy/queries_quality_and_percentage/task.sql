select t1.query_name, t1.quality, ifnull(t2.poor_query_percentage, 0) as poor_query_percentage
from (
         select
             query_name,
             round(sum((rating/position)) / count(query_name), 2)  as quality
         from Queries
         group by query_name
     ) t1 left join (
    select q1.query_name, round((q1.cnt / q2.cnt) * 100, 2) as poor_query_percentage
    from (
             select query_name, count(*) as cnt
             from Queries
             where rating < 3
             group by query_name
         ) q1 left join (
        select query_name, count(*) as cnt
        from Queries
        group by query_name
    ) q2 using(query_name)
) t2 using (query_name)
where query_name is not null;
