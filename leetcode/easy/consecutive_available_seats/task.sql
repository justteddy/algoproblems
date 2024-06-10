# Write your MySQL query statement below
select x.seat_id
from(
        select
            seat_id,
            free,
            lead(free, 1) over(order by seat_id) as next_seat_free,
                lag(free, 1) over(order by seat_id) as prev_seat_free
        from Cinema
        order by seat_id asc
    ) x
where x.free = 1 and (x.next_seat_free = 1 or x.prev_seat_free = 1)
order by seat_id;
