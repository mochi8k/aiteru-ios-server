# Getting Started

## MySQL
```
$ brew install mysql
$ mysql.server start
$ mysql -uroot

mysql> create database aiteru;

mysql> create table users(
    `id` int(11) unsigned not null auto_increment,
    `user_name` varchar(255) not null,
    `created_at` date,
    `created_by` int(11),
    `updated_at` date,
    `updated_by` int(11),
    primary key(id)
);

mysql> create table places(
    `id` int(11) unsigned not null auto_increment,
    `place_name` varchar(255) not null,
    `owners` varchar(255) not null,
    `collaborators` varchar(255) not null,
    `created_at` date,
    `created_by` int(11),
    `updated_at` date,
    `updated_by` int(11),
     primary key(id)
);

mysql> create table place_status(
   `place_id` int(11) unsigned not null,
   `is_open` boolean not null,
   `updated_at` date,
   `updated_by` int(11)
);

```