create table listings (
    id UUID primary key default gen_random_uuid(),
    title text not null,
    description text not null,
    price bigint not null,
    city text not null,
    created_at timestampz not null default now ()
);