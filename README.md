# Getting Started

## MySQL
```
$ brew install mysql
$ mysql.server start
$ mysql -uroot

create database aiteru;
```

# Progress

POST: /auth
```
{
  "name": "user-name",
  "address": "e-mail"
}
```
GET: /v1/places  
GET: /v1/places/{place-id}