# Getting Started

## MySQL
```
$ brew install mysql
$ mysql.server start
$ mysql -uroot

create database aiteru;

create table users(
    `id` int(11) unsigned not null auto_increment,
    `user_name` varchar(255) not null,
    `created_at` timestamp not null default current_timestamp,
    `created_by` int(11),
    `updated_at` timestamp not null default current_timestamp on update current_timestamp,
    `updated_by` int(11),
    primary key(id)
);

create table places(
    `id` int(11) unsigned not null auto_increment,
    `place_name` varchar(255) not null,
    `created_at` timestamp not null default current_timestamp,
    `created_by` int(11),
    `updated_at` timestamp not null default current_timestamp on update current_timestamp,
    `updated_by` int(11),
    primary key(id)
);

create table place_owners(
    `place_id` int(11) unsigned not null,
    `owner_id` int(11) unsigned not null,
    primary key(place_id, owner_id)
);

create table place_collaborators(
    `place_id` int(11) unsigned not null,
    `collaborator_id` int(11) unsigned not null,
    primary key(place_id, collaborator_id)
);

create table place_status(
    `place_id` int(11) unsigned not null,
    `is_open` boolean not null,
    `updated_at` timestamp not null default current_timestamp,
    `updated_by` int(11)
);

```