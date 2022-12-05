CREATE TABLE users (
   id BIGSERIAL primary key,
   name TEXT not null,
   last_name TEXT,
   age INTEGER,
   created_at TIMESTAMP default now()
);