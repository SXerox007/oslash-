create table all_tweets (
id uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
user_id uuid REFERENCES all_users(id),
Tweet text not null,
state integer default 1,
version text,
capture_time timestamp default current_timestamp
);