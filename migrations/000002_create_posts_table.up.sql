create table if not exists posts(
    id serial primary key,
    user_id integer references users(id) on delete cascade,
    title varchar(200) not null,
    content text,
    created_at timestamp default current_timestamp
);