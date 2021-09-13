create table tasks
(
    id uuid primary key,
    name varchar (255) not null,
    description text,
    created_at timestamp not null default current_timestamp,
    status varchar(255) default 'Not Started',
    due_date timestamp
);

drop table tasks
