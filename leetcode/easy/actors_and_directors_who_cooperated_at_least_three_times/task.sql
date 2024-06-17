select x.actor_id, x.director_id
from(
        select actor_id, director_id, count(*) as total
        from ActorDirector
        group by actor_id, director_id
        having total >= 3
    ) as x
