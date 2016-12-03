create table users(
    `id` int(11) unsigned not null auto_increment,
    `user_name` varchar(255) not null,
    `created_at` timestamp not null default current_timestamp,
    `created_by` int(11),
    `updated_at` timestamp not null default current_timestamp on update current_timestamp,
    `updated_by` int(11),
    primary key(id)
);
