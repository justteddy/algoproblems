select t.player_id, t2.player_name, count(t.player_id) as grand_slams_count
from(
        select Wimbledon as player_id
        from Championships
        UNION ALL
        select Fr_open as player_id
        from Championships
        UNION ALL
        select US_open as player_id
        from Championships
        UNION ALL
        select Au_open as player_id
        from Championships
    ) as t left join Players t2 using(player_id)
group by t.player_id
having count(t.player_id) > 0
