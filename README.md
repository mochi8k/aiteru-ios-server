# Getting Started

## MySQL
```
$ brew install mysql
$ mysql.server start
$ mysql -uroot

mysql> create database gosample;
mysql> use gosample;
mysql> create table attendance (id int not null auto_increment, username varchar(64), time DATETIME, primary key (id));
```