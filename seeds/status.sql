select pid, state, query_start, query
from pg_stat_activity
where
    query like '%pg_sleep%'
    and state != 'idle';
