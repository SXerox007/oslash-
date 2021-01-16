create table all_users (
id uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
phone text,
first_name text,
last_name text,
password text not null,
email text not null,
role text,
version text,
capture_time timestamp default current_timestamp
);