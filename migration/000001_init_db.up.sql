CREATE TABLE links
(
    id serial primary key,
    long_link varchar not null ,
    short_link varchar(10) unique not null
)