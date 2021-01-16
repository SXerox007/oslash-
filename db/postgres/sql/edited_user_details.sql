create table edited_user_details (
id uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
phone text,
first_name text,
last_name text,
email text not null,
role text,
user_id uuid REFERENCES all_users(id),
admin_id uuid REFERENCES all_users(id),
approved boolean default false,
approver_by uuid REFERENCES all_users(id),
version text,
capture_time timestamp default current_timestamp
);
