create table users (
   id         serial primary key,
   name       text not null,
   email      text not null unique,
   password   text not null,
   created_at timestamp default now()
);