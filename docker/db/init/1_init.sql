create table users (
  id serial primary key,
  email text not null unique,
  password text not null
);

insert into users (email, password) values ('first@sample.co.jp', 'golang');