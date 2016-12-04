create table users(
    `id` int(11) unsigned not null auto_increment,
    `user_name` varchar(255) not null,
    `created_at` datetime,
    `created_by` int(11),
    `updated_at` datetime,
    `updated_by` int(11),
    primary key(id)
);
