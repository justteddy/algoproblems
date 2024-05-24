# Write your MySQL query statement below
select t3.machine_id, round(sum(t3.timestamp_end-t3.timestamp_start) / count(distinct t3.process_id), 3) as processing_time
from(
        select t1.machine_id, t1.process_id, t1.timestamp as timestamp_start, t2.timestamp as timestamp_end
        from (
                 select *
                 from Activity
                 where activity_type = 'start'
             ) t1 left join
             (
                 select *
                 from Activity
                 where activity_type = 'end'
             ) t2 on t1.machine_id = t2.machine_id and t1.process_id = t2.process_id
        GROUP BY t1.machine_id, t1.process_id
    ) t3
group by machine_id;
